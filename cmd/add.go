package cmd

import (
	. "github.com/xiabai84/githooks/hooks"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new workspace with githooks",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		AddGithooks()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
