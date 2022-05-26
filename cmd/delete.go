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
	Short: "Delete a githooks project and its settings",
	Long:  `A longer description `,
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
			name := strings.Replace(strings.ToLower(hook.Project), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
			return strings.Contains(name, input)
		}

		prompt1 := promptui.Select{
			Label:     "Delete:",
			Items:     hookArr,
			Templates: templates,
			Size:      5,
			Searcher:  searcher,
		}

		i, _, err := prompt1.Run()
		CheckError(err)

		if hookArr[i].Project != "Quit" {
			prompt2 := promptui.Prompt{
				Label:     "Really want to delete this project",
				IsConfirm: true,
			}
			confirmed, err := prompt2.Run()
			if err != nil {
				fmt.Println("Canceled")
			}
			if confirmed != "y" {
				os.Exit(1)
			}
			DeleteSelectedProject(&hookArr[i])
		} else {
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
