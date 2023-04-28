package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/emicklei/proto"
)

func main() {
	var startDir string
	var outputBaseDir string

	flag.StringVar(&startDir, "input", "./proto", "Input directory to start searching for proto files")
	flag.StringVar(&outputBaseDir, "output", "./docs/spec", "Output base directory for generated markdown files")
	flag.Parse()

	err := filepath.Walk(startDir, func(path string, f os.FileInfo, err error) error {
		return visit(path, f, err, outputBaseDir)
	})
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", startDir, err)
	}
}

func visit(path string, f os.FileInfo, err error, outputBaseDir string) error {
	if err != nil {
		fmt.Printf("Error visiting %q: %v\n", path, err)
		return err
	}

	if f.IsDir() {
		return nil
	}

	if filepath.Ext(path) == ".proto" {
		processProtoFile(path, outputBaseDir)
	}

	return nil
}

func processProtoFile(path string, outputBaseDir string) {
	reader, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening proto file %q: %v\n", path, err)
		return
	}
	defer reader.Close()

	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		fmt.Printf("Error parsing proto file %q: %v\n", path, err)
		return
	}

	var packageName string
	var msgServices []*proto.Service
	messageMap := make(map[string]*proto.Message)

	proto.Walk(definition,
		proto.WithPackage(func(p *proto.Package) {
			packageName = p.Name
		}),
		proto.WithService(func(s *proto.Service) {
			if s.Name == "Msg" {
				msgServices = append(msgServices, s)
			}
		}),
		proto.WithMessage(func(m *proto.Message) {
			messageMap[m.Name] = m
		}),
	)

	if packageName != "" && len(msgServices) > 0 {
		outputDir := filepath.Join(outputBaseDir, getLastSegmentOfPackageName(packageName))
		err = os.MkdirAll(outputDir, 0755)
		if err != nil {
			fmt.Printf("Error creating directory %q: %v\n", outputDir, err)
			return
		}

		outputFile := filepath.Join(outputDir, "messages.md")
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("Error creating file %q: %v\n", outputFile, err)
			return
		}
		defer file.Close()

		file.WriteString("# Messages\n\n")
		for _, service := range msgServices {
			for _, element := range service.Elements {
				if rpc, ok := element.(*proto.RPC); ok {
					functionName := strings.TrimPrefix(rpc.RequestType, "Msg")
					if message, found := messageMap[rpc.RequestType]; found {
						// Get the current working directory
						currentDir, err := os.Getwd()
						if err != nil {
							fmt.Printf("Error getting current working directory: %v\n", err)
							return
						}

						// Search for the corresponding Go function
						goFunctionPath := filepath.Join(currentDir, "x")
						functionComment, functionFound := findFunctionInGoFiles(functionName, goFunctionPath)

						file.WriteString(fmt.Sprintf("## %s\n\n", rpc.RequestType))
						if functionFound && functionComment != "" {
							file.WriteString(fmt.Sprintf("%s\n", functionComment))
						}
						file.WriteString("```proto\n")
						file.WriteString(messageToString(message))
						file.WriteString("```\n\n")
					} else {
						file.WriteString(fmt.Sprintf("## %s\n\n```\n%s\n```\n\n", rpc.RequestType, rpc.RequestType))
					}
				}
			}
		}
	}
}

func messageToString(message *proto.Message) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("message %s {\n", message.Name))

	for _, element := range message.Elements {
		switch field := element.(type) {
		case *proto.NormalField:
			sb.WriteString(fmt.Sprintf("\t%s %s = %d;\n", field.Type, field.Name, field.Sequence))
		case *proto.MapField:
			sb.WriteString(fmt.Sprintf("\tmap<%s, %s> %s = %d;\n", field.KeyType, field.Type, field.Name, field.Sequence))
		case *proto.Oneof:
			sb.WriteString(fmt.Sprintf("\toneof %s {\n", field.Name))
			for _, of := range field.Elements {
				if oneOfField, ok := of.(*proto.OneOfField); ok {
					sb.WriteString(fmt.Sprintf("\t\t%s %s = %d;\n", oneOfField.Type, oneOfField.Name, oneOfField.Sequence))
				}
			}
			sb.WriteString("\t}\n")
		}
	}

	sb.WriteString("}\n")

	return sb.String()
}

func findFunctionInGoFiles(functionName string, startDir string) (string, bool) {
	var functionDoc string
	found := false

	fileSet := token.NewFileSet()
	err := filepath.Walk(startDir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if f.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}

		node, err := parser.ParseFile(fileSet, path, nil, parser.ParseComments)
		if err != nil {
			fmt.Printf("Error parsing Go file %q: %v\n", path, err)
			return err
		}

		for _, decl := range node.Decls {
			if funcDecl, ok := decl.(*ast.FuncDecl); ok {
				if funcDecl.Name.Name == functionName {
					if funcDecl.Doc != nil {
						functionDoc = funcDecl.Doc.Text()
					}
					found = true
					return filepath.SkipDir
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", startDir, err)
	}

	return functionDoc, found
}

func getLastSegmentOfPackageName(packageName string) string {
	segments := strings.Split(packageName, ".")
	return segments[len(segments)-1]
}
