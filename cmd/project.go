package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gitzup/agent/pkg"
	"github.com/gitzup/gcp/internal"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Applies GCP project resources.",
}

var projectInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a GCP project resource.",
	Run: func(cmd *cobra.Command, args []string) {

		projectSchemaBytes, err := internal.Asset("schema/project.json")
		if err != nil {
			panic(err)
		}

		var projectSchema interface{}
		err = json.Unmarshal(projectSchemaBytes, &projectSchema)
		if err != nil {
			panic(err)
		}

		response := pkg.ResourceInitResponse{
			ConfigSchema: projectSchema,
			DiscoveryAction: pkg.Action{
				Name:       "discover",
				Image:      "gitzup/gcp",
				Entrypoint: []string{"/app/agent"},
				Cmd:        []string{"discover"},
			},
		}
		json, err := json.Marshal(&response)
		if err != nil {
			panic(err)
		}

		err = os.MkdirAll("/gitzup", 0755)
		if err != nil {
			panic(err)
		}

		ioutil.WriteFile("/gitzup/result.json", json, 0644)
		fmt.Println("project called")
	},
}

var projectDiscoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Discovers the state of the GCP project.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project called")
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectInitCmd)
	projectCmd.AddCommand(projectDiscoverCmd)
}
