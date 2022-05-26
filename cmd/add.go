package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	. "github.com/stefan-niemeyer/githooks/hooks"
	. "github.com/stefan-niemeyer/githooks/types"
	. "github.com/stefan-niemeyer/githooks/utils"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new workspace with githooks",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {

		CheckConfigFiles()

		projName := GetPromptInput(Dialog{
			ErrorMsg: "Please provide a Jira project key to track.",
			Label:    "Enter your Jira project key:",
		}, "")

		cwd, errCwd := os.Getwd()
		CheckError(errCwd)
		homeDir := os.Getenv("HOME")
		if len(homeDir) != 0 {
			cwd = strings.Replace(cwd, homeDir, "~", 1)
		}
		if !strings.HasSuffix(cwd, "/") {
			cwd += "/"
		}
		workDir := GetPromptInput(Dialog{
			ErrorMsg: "Please enter a path to your workspace.",
			Label:    fmt.Sprintf("Enter path to your workspace (%s):", cwd),
		}, cwd)

		if !strings.HasSuffix(workDir, "/") {
			workDir += "/"
		}

		newHook := GitHooks{Project: projName, JiraName: strings.ToUpper(projName), WorkDir: workDir}
		PreviewConfig(&newHook)

		prompt := promptui.Prompt{
			Label:     "Input was correct",
			IsConfirm: true,
		}

		confirmed, err := prompt.Run()
		if err != nil {
			fmt.Println(promptui.IconBad + " Cancelled adding of a new githooks project.")
		}

		if confirmed == "y" {
			AddGithooks(&newHook)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
