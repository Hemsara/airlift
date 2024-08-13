/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	database "airlift/internal/connections"
	"airlift/internal/styles"
	"airlift/schemas"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var UploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "🦦 Build and upload testflight binaries",

	Run: func(cmd *cobra.Command, args []string) {
		q, _ := cmd.Flags().GetString("project")

		var project schemas.Project
		if err := database.DB.Where("project_name = ?", q).First(&project).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Println("Project not found")
			} else {
				log.Fatalf("failed to get project: %v", err)
			}
		}

		fmt.Println(styles.SuccessStyle.Render("Project found, Building binary.."))
		ipaCmd := exec.Command("flutter", "build", "ipa")
		ipaCmd.Dir = project.Path

		ipaCmd.Stdout = os.Stdout
		ipaCmd.Stderr = os.Stderr

		err := ipaCmd.Run()
		if err != nil {
			fmt.Printf("Error running flutter build ipa: %v\n", err)
			return
		}
		fmt.Println(styles.SuccessStyle.Render("IPA build completed, now uploading with xcrun..."))
		ipaDirPath := filepath.Join(project.Path, "build", "ios", "ipa")

		files, err := os.ReadDir(ipaDirPath)
		if err != nil {
			fmt.Printf("Error reading directory: %v\n", err)
			return
		}
		var ipaFiles []string

		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".ipa") {
				ipaFilePath := filepath.Join(ipaDirPath, file.Name())
				ipaFiles = append(ipaFiles, ipaFilePath)
			}
		}

		if len(ipaFiles) == 0 {
			fmt.Println("No .ipa files found in the directory.")
			return
		}
		ipa := ipaFiles[0]

		altoolCmd := exec.Command("xcrun", "altool", "--upload-app", "--type", "ios", "-f", ipa, "--apiKey", project.KeyID, "--apiIssuer", project.IssueID)
		altoolCmd.Stdout = os.Stdout
		altoolCmd.Stderr = os.Stderr

		err = altoolCmd.Run()
		if err != nil {
			fmt.Printf("Error running xcrun altool: %v\n", err)
			return
		}

		fmt.Println(styles.SuccessStyle.Render("Upload completed successfully!"))

	},
}

func init() {
	UploadCmd.Flags().StringP("project", "p", "", "Project name")
	UploadCmd.MarkFlagRequired("project")

}
