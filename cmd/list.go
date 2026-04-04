package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	. "github.com/stefan-niemeyer/githooks/hooks"
	. "github.com/stefan-niemeyer/githooks/prompt"
	. "github.com/stefan-niemeyer/githooks/styles"
	. "github.com/stefan-niemeyer/githooks/types"
	. "github.com/stefan-niemeyer/githooks/utils"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all workspaces managed by githooks",
	Long:  `List all workspaces managed by githooks`,
	Run: func(cmd *cobra.Command, args []string) {
		CheckConfigFiles()

		ghConfig := ReadGitHooksConfig()
		empty := Workspace{Name: "Quit"}
		preselectIdx := GetWorkspaceIndex(ghConfig.Workspaces)
		workspaces := append(ghConfig.Workspaces, empty)

		prompt := promptui.Select{
			Label:     "Active githooks workspaces:",
			Items:     workspaces,
			Templates: GetDefaultSelectTemplates(),
			Size:      5,
			Searcher:  NewWorkspaceSearcher(workspaces),
			CursorPos: preselectIdx,
		}

		_, _, err := prompt.Run()
		CheckError(err)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
