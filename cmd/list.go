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
	Short: "List all from githooks managed projects.",
	Long:  `List all githooks managed projects`,
	Run: func(cmd *cobra.Command, args []string) {
		hookArr := ReadFromGitHookLog()
		empty := GitHooks{Project: "Quit"}
		hookArr = append(hookArr, empty)

		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "➣ {{ .Project | cyan }}",
			Inactive: "  {{ .Project | cyan }}",
			Selected: "➣ {{ .Project | red | cyan }}",
			Details:  DetailTmpl,
		}

		searcher := func(input string, index int) bool {
			hook := hookArr[index]
			projName := strings.Replace(strings.ToLower(hook.Project), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
			return strings.Contains(projName, input)
		}

		prompt := promptui.Select{
			Label:     "Active Githooks config file(s):",
			Items:     hookArr,
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
	listCmd.PersistentFlags().StringP("list", "l", "", "List all githooks projects")
}
