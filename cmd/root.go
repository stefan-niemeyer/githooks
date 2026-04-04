package cmd

import (
	"github.com/stefan-niemeyer/githooks/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "githooks",
	Short: "githooks helps developers with setting name conventions of a Git commit message",
	Long: `githooks ensures the existence of Jira issue keys in commit messages.
Compliance with conventional commits can also be ensured.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.CheckArgs(cmd, args)
	},
}

func Execute() {
	err := rootCmd.Execute()
	utils.CheckError(err)
}
