package alter

import (
	database "airlift/internal/connections"
	"airlift/internal/styles"
	"airlift/schemas"
	"fmt"

	"github.com/spf13/cobra"
)

var AlterCmd = &cobra.Command{
	Use:   "alter",
	Short: "🧟 Edit or delete your existing ",

	Run: func(cmd *cobra.Command, args []string) {
		projectID, _ := cmd.Flags().GetInt("id")

		title, _ := cmd.Flags().GetString("title")

		issuerID, _ := cmd.Flags().GetString("issuer_id")
		path, _ := cmd.Flags().GetString("path")
		keyID, _ := cmd.Flags().GetString("key_id")

		var project schemas.Project

		if err := database.DB.First(&project, projectID).Error; err != nil {
			fmt.Println(styles.ErrStyle.Render("Project not found"))
		}
		if title != "" {
			project.ProjectName = title
		}
		if issuerID != "" {
			project.IssueID = issuerID
		}
		if path != "" {
			project.Path = path
		}
		if keyID != "" {
			project.KeyID = keyID

		}
		if err := database.DB.Save(&project).Error; err != nil {
			fmt.Println(styles.ErrStyle.Render("Unable to save changes"))

		}

		fmt.Println(styles.SuccessStyle.Render("Project updated!"))
	},
}

func init() {
	AlterCmd.Flags().IntP("id", "i", 0, "Project ID to edit")
	AlterCmd.Flags().StringP("title", "t", "", "New title for the project")
	AlterCmd.Flags().StringP("issuer_id", "s", "", "New issuer ID for the project")
	AlterCmd.Flags().StringP("path", "p", "", "New path for the project")
	AlterCmd.Flags().StringP("key_id", "k", "", "New key ID for the project")

	AlterCmd.MarkFlagRequired("id")

}
