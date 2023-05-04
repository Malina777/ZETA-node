package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zeta-chain/zetacore/app"
	"os"
	"path/filepath"
)

func CollectObserverInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collect-observer-info [folder]",
		Short: "collect observer info from a folder , default path is ~/.zetacored/os_info/ \n",
		Args:  cobra.MaximumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			defaultHome := app.DefaultNodeHome
			defaultFile := filepath.Join(defaultHome, "os_info")
			if len(args) == 0 {
				args = append(args, defaultFile)
			}
			directory := args[0]
			files, err := os.ReadDir(directory)
			if err != nil {
				return err
			}
			var observerInfoList []ObserverInfoReader
			err = os.Chdir(directory)
			if err != nil {
				return err
			}
			for _, file := range files {
				var observerInfo ObserverInfoReader
				info, err := file.Info()
				if err != nil {
					return err
				}
				f, err := os.ReadFile(info.Name())
				if err != nil {
					return err
				}
				err = json.Unmarshal(f, &observerInfo)
				if err != nil {
					return err
				}
				observerInfoList = append(observerInfoList, observerInfo)
			}
			fmt.Println(observerInfoList)
			file, _ := json.MarshalIndent(observerInfoList, "", " ")
			_ = os.WriteFile("observer_info.json", file, 0600)
			return nil
		},
	}
	return cmd
}
