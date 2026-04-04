package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	. "github.com/stefan-niemeyer/githooks/hooks"
	. "github.com/stefan-niemeyer/githooks/prompt"
	. "github.com/stefan-niemeyer/githooks/styles"
	. "github.com/stefan-niemeyer/githooks/types"
	. "github.com/stefan-niemeyer/githooks/utils"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a githooks workspace and its settings",
	Long:  `Delete a githooks workspace and its settings`,
	Run: func(cmd *cobra.Command, args []string) {
		CheckConfigFiles()

		ghConfig := ReadGitHooksConfig()
		empty := Workspace{Name: "Quit"}
		preselectIdx := GetWorkspaceIndex(ghConfig.Workspaces)
		workspaces := append(ghConfig.Workspaces, empty)

		prompt1 := promptui.Select{
			Label:     "Delete:",
			Items:     workspaces,
			Templates: GetDefaultSelectTemplates(),
			Size:      5,
			Searcher:  NewWorkspaceSearcher(workspaces),
			CursorPos: preselectIdx,
		}

		i, _, err := prompt1.Run()
		CheckError(err)

		if workspaces[i].Name != "Quit" {
			prompt2 := promptui.Prompt{
				Label:     "Do you Really want to delete this workspace",
				IsConfirm: true,
			}
			confirmed, err := prompt2.Run()
			if err != nil {
				fmt.Println("Canceled")
			}
			if strings.ToLower(confirmed) != "y" {
				os.Exit(1)
			}
			DeleteSelectedWorkspace(&ghConfig, i)
		} else {
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
