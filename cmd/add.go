package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	. "github.com/stefan-niemeyer/githooks/hooks"
	. "github.com/stefan-niemeyer/githooks/prompt"
	. "github.com/stefan-niemeyer/githooks/styles"
	. "github.com/stefan-niemeyer/githooks/utils"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new workspace with githooks",
	Long:  `Add a new workspace with githooks`,
	Run: func(cmd *cobra.Command, args []string) {

		CheckConfigFiles()

		cwd, errCwd := os.Getwd()
		CheckError(errCwd)
		newWorkspace := CreateWorkspace("", "", cwd, "", PlainBrackets)
		PreviewConfig(&newWorkspace)

		prompt := promptui.Prompt{
			Label:     "Input was correct",
			IsConfirm: true,
		}

		confirmed, err := prompt.Run()
		if err != nil {
			fmt.Println(promptui.IconBad + " Canceled adding of a new githooks workspace")
		}

		if strings.ToLower(confirmed) == "y" {
			AddWorkspace(&newWorkspace)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
