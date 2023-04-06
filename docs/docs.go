package docs

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	apiFile   = "/static/openapi.yml"
	indexFile = "template/index.tpl"
)

//go:embed static
var staticFS embed.FS

//go:embed template
var templateFS embed.FS

func RegisterOpenAPIService(router *mux.Router) {
	router.Handle(apiFile, http.FileServer(http.FS(staticFS)))
	router.HandleFunc("/", openAPIHandler())
}

func openAPIHandler() http.HandlerFunc {
	tmpl, _ := template.ParseFS(templateFS, indexFile)

	return func(w http.ResponseWriter, req *http.Request) {
		tmpl.Execute(w, struct {
			URL string
		}{
			apiFile,
		})
	}
}
