package cmd

import (
	"github.com/spf13/cobra"
	. "github.com/stefan-niemeyer/githooks/hooks"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup your githooks workspace configuration",
	Long:  `Setup your githooks workspace configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		InitHooks()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
