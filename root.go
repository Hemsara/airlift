package main

import (
	add "airlift/cmd/add"
	alter "airlift/cmd/alter"
	database "airlift/internal/connections"
	"airlift/pkg/initializers"

	list "airlift/cmd/list"
	upload "airlift/cmd/upload"

	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "airlift",
	Short: "🧪 Easily manage of testflight releases",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommandPallet() {
	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(alter.AlterCmd)
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(upload.UploadCmd)

}
func init() {
	initializers.LoadENV()
	database.New()
	initializers.MakeMigrations()

	addSubCommandPallet()
}

func main() {
	Execute()
}
