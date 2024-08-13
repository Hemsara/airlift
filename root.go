package main

import (
	build "airlift/cmd/build"

	initialize "airlift/cmd/init"

	list "airlift/cmd/list"
	upload "airlift/cmd/upload"

	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "airlift",
	Short: "A brief description of your application",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommandPallet() {
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(initialize.InitCmd)
	rootCmd.AddCommand(build.BuildCmd)

	rootCmd.AddCommand(upload.UploadCmd)

}
func init() {

	addSubCommandPallet()
}

func main() {
	Execute()
}
