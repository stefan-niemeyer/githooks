package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	. "github.com/stefan-niemeyer/githooks/config"
	. "github.com/stefan-niemeyer/githooks/hooks"
	. "github.com/stefan-niemeyer/githooks/types"
	. "github.com/stefan-niemeyer/githooks/utils"
	"strings"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all workspaces managed by githooks.",
	Long:  `List all workspaces managed by githooks`,
	Run: func(cmd *cobra.Command, args []string) {
		CheckConfigFiles()

		ghConfig := ReadGitHooksConfig()
		empty := Workspace{Name: "Quit"}
		ghConfig.Workspaces = append(ghConfig.Workspaces, empty)

		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "➣ {{ .Name | cyan }}",
			Inactive: "  {{ .Name | cyan }}",
			Selected: "➣ {{ .Name | red | cyan }}",
			Details:  DetailTmpl,
		}

		searcher := func(input string, index int) bool {
			workspace := ghConfig.Workspaces[index]
			name := strings.Replace(strings.ToLower(workspace.Name), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
			return strings.Contains(name, input)
		}

		prompt := promptui.Select{
			Label:     "Active githooks workspaces:",
			Items:     ghConfig.Workspaces,
			Templates: templates,
			Size:      5,
			Searcher:  searcher,
		}

		_, _, err := prompt.Run()
		CheckError(err)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
