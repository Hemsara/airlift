package add

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "🔥 Register a new app to airlift",
	Run: func(cmd *cobra.Command, args []string) {
		project, _ := cmd.Flags().GetString("project")
		path, _ := cmd.Flags().GetString("path")
		apiKey, _ := cmd.Flags().GetString("apiKey")
		apiIssuer, _ := cmd.Flags().GetString("apiIssuer")

		fmt.Printf("Project: %s\n", project)
		fmt.Printf("Path: %s\n", path)
		fmt.Printf("API Key: %s\n", apiKey)
		fmt.Printf("API Issuer: %s\n", apiIssuer)

	},
}

func init() {
	AddCmd.Flags().StringP("project", "p", "", "Project name")
	AddCmd.Flags().StringP("path", "P", "", "Path to the app")
	AddCmd.Flags().StringP("apiKey", "k", "", "API key for authentication")
	AddCmd.Flags().StringP("apiIssuer", "i", "", "API issuer for authentication")

	AddCmd.MarkFlagRequired("project")
	AddCmd.MarkFlagRequired("apiKey")
	AddCmd.MarkFlagRequired("path")
	AddCmd.MarkFlagRequired("apiIssuer")

}
