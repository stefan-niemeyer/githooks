package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	. "github.com/stefan-niemeyer/githooks/config"
	. "github.com/stefan-niemeyer/githooks/hooks"
	. "github.com/stefan-niemeyer/githooks/types"
	. "github.com/stefan-niemeyer/githooks/utils"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a githooks workspace and its settings.",
	Long:  `Delete a githooks workspace and its settings`,
	Run: func(cmd *cobra.Command, args []string) {
		CheckConfigFiles()

		ghConfig := ReadGitHooksConfig()
		empty := Workspace{Name: "Quit"}
		workspaces := append(ghConfig.Workspaces, empty)

		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "➣ {{ .Name | cyan }}",
			Inactive: "  {{ .Name | cyan }}",
			Selected: "➣ {{ .Name | red | cyan }}",
			Details:  DetailTmpl,
		}

		searcher := func(input string, index int) bool {
			workspace := workspaces[index]
			name := strings.Replace(strings.ToLower(workspace.Name), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
			return strings.Contains(name, input)
		}

		prompt1 := promptui.Select{
			Label:     "Delete:",
			Items:     workspaces,
			Templates: templates,
			Size:      5,
			Searcher:  searcher,
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
