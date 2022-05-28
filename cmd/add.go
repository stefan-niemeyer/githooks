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
	Long:  `Add a new workspace with githooks`,
	Run: func(cmd *cobra.Command, args []string) {

		CheckConfigFiles()

		projName := GetPromptInput(Dialog{
			ErrorMsg: "Please provide a name for the workspace.",
			Label:    "Enter your workspace name:",
		}, "")

		jiraName := strings.ToUpper(projName)
		jiraName = GetPromptInput(Dialog{
			ErrorMsg: "Please provide a Jira project key RegEx to track, e.g. ALPHA or (ALPHA|BETA)",
			Label:    fmt.Sprintf("Enter your Jira project key RegEx (%s):", jiraName),
		}, jiraName)

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

		newWorkspace := Workspace{
			Name:         projName,
			ProjectKeyRE: strings.ToUpper(jiraName),
			Folder:       workDir,
		}
		PreviewConfig(&newWorkspace)

		prompt := promptui.Prompt{
			Label:     "Input was correct",
			IsConfirm: true,
		}

		confirmed, err := prompt.Run()
		if err != nil {
			fmt.Println(promptui.IconBad + " Canceled adding of a new githooks workspace.")
		}

		if strings.ToLower(confirmed) == "y" {
			AddWorkspace(&newWorkspace)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
