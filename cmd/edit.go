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

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a githooks workspace and its settings",
	Long:  `Edit a githooks workspace and its settings`,
	Run: func(cmd *cobra.Command, args []string) {
		CheckConfigFiles()

		ghConfig := ReadGitHooksConfig()
		empty := Workspace{Name: "Quit"}
		preselectIdx := GetWorkspaceIndex(ghConfig.Workspaces)
		workspaces := append(ghConfig.Workspaces, empty)

		prompt1 := promptui.Select{
			Label:     "Edit:",
			Items:     workspaces,
			Templates: GetDefaultSelectTemplates(),
			Size:      5,
			Searcher:  NewWorkspaceSearcher(workspaces),
			CursorPos: preselectIdx,
		}

		i, _, err := prompt1.Run()
		currentWs := workspaces[i]
		CheckError(err)

		if currentWs.Name != "Quit" {
			newWs := CreateWorkspace(currentWs.Name, currentWs.ProjectKeyRE, currentWs.Folder, currentWs.AllowedTypes, FormatStyle(currentWs.CommitMessageStyle))

			numChanges := 0
			if currentWs.Name != newWs.Name {
				fmt.Printf("Workspace name changed\n from: %s\n   to: %s\n", currentWs.Name, newWs.Name)
				numChanges++
			}
			if currentWs.ProjectKeyRE != newWs.ProjectKeyRE {
				fmt.Printf("Project key RE changed\n from: %s\n   to: %s\n", currentWs.ProjectKeyRE, newWs.ProjectKeyRE)
				numChanges++
			}
			if currentWs.CommitMessageStyle != newWs.CommitMessageStyle {
				fmt.Printf("Format style changed\n from: %s\n   to: %s\n", currentWs.CommitMessageStyle, newWs.CommitMessageStyle)
				numChanges++
			}
			if currentWs.AllowedTypes != newWs.AllowedTypes {
				fmt.Printf("Allowed types changed\n from: %s\n   to: %s\n", currentWs.AllowedTypes, newWs.AllowedTypes)
				numChanges++
			}
			if currentWs.Folder != newWs.Folder {
				fmt.Printf("Folder changed\n from: %s\n   to: %s\n", currentWs.Folder, newWs.Folder)
				numChanges++
			}
			if numChanges == 0 {
				fmt.Println(promptui.IconWarn + " No changes were detected.")
				os.Exit(0)
			}

			prompt := promptui.Prompt{
				Label:     "Changes are correct",
				IsConfirm: true,
			}

			confirmed, err := prompt.Run()
			if err != nil {
				fmt.Println(promptui.IconBad + " Canceled editing of a githooks workspace")
			}
			if strings.ToLower(confirmed) == "y" {
				DeleteSelectedWorkspace(&ghConfig, i)
				AddWorkspace(&newWs)
			}
		} else {
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
