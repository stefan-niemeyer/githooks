package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	. "github.com/stefan-niemeyer/githooks/hooks"
	. "github.com/stefan-niemeyer/githooks/types"
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
			ErrorMsg: "❌ Please provide a Jira project key to track.",
			Label:    "🌟 Enter your Jira project's name:",
		})

		workDir := GetPromptInput(Dialog{
			ErrorMsg: "❌ Please a path to workspace.",
			Label:    fmt.Sprintf("🌟 Enter path to your workspace:"),
		})

		hasSlash := strings.HasSuffix(workDir, "/")
		if !hasSlash {
			workDir = workDir + "/"
		}

		newHook := GitHooks{Project: projName, JiraName: strings.ToUpper(projName), WorkDir: workDir}
		PreviewConfig(&newHook)

		prompt := promptui.Prompt{
			Label:     "Input was correct",
			IsConfirm: true,
		}

		confirmed, err := prompt.Run()
		if err != nil {
			fmt.Println("❌ Canceled setting new githooks project.")
		}

		if confirmed == "y" {
			AddGithooks(&newHook)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
