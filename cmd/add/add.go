package add

import (
	database "airlift/internal/connections"
	"airlift/internal/styles"
	"airlift/schemas"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

		expandedPath, err := ExpandUserPath(path)
		if err != nil {
			fmt.Println(styles.ErrStyle.Render("Error getting file info: " + err.Error()))
			return
		}

		if _, err := os.Stat(expandedPath); err != nil {
			fmt.Println(styles.ErrStyle.Render("Error accessing file: " + err.Error()))
			return
		}

		abs, err := filepath.Abs(expandedPath)
		if err != nil {
			fmt.Println(styles.ErrStyle.Render("Error getting absolute path: " + err.Error()))
			return
		}

		db := database.DB

		session := schemas.Project{
			IssueID:     apiIssuer,
			KeyID:       apiKey,
			ProjectName: project,
			Path:        abs,
		}
		if err := db.Create(&session).Error; err != nil {
			fmt.Println(styles.ErrStyle.Render("Error: Unable to add project. " + err.Error()))
			return
		}

		fmt.Println(styles.SuccessStyle.Render("🎉 Project registration successful! 🚀"))
	},
}

func init() {
	AddCmd.Flags().StringP("project", "p", "", "Project name")
	AddCmd.Flags().StringP("path", "P", "", "Path to your flutter app")
	AddCmd.Flags().StringP("apiKey", "k", "", "Apple API key ID")
	AddCmd.Flags().StringP("apiIssuer", "i", "", "Apple API issuer ID")

	AddCmd.MarkFlagRequired("project")
	AddCmd.MarkFlagRequired("apiKey")
	AddCmd.MarkFlagRequired("path")
	AddCmd.MarkFlagRequired("apiIssuer")
}

func ExpandUserPath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("unable to get home directory: %w", err)
		}
		path = filepath.Join(homeDir, path[2:])
	}
	return path, nil
}
