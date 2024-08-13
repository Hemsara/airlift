package list

import (
	database "airlift/internal/connections"
	"airlift/internal/styles"
	"airlift/schemas"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List out registered apps",
	Run: func(cmd *cobra.Command, args []string) {

		var projects []schemas.Project
		result := database.DB.Find(&projects)

		if result.Error != nil {
			fmt.Println(styles.ErrStyle.Render("Error Fetching your projects"))
			fmt.Println(styles.SuccessStyle.Render("Your projects"))
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

		for _, project := range projects {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t\n",
				project.ID, project.ProjectName, project.IssueID, project.KeyID)
		}

		w.Flush()

	},
}

func init() {

}
