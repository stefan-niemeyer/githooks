package cmd

import (
	"github.com/spf13/cobra"
	. "github.com/xiabai84/githooks/hooks"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "setup your local githooks",
	Long:  `setup your local githooks`,
	Run: func(cmd *cobra.Command, args []string) {
		InitHooks()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.PersistentFlags().StringP("init", "i", "", "setup your githooks")
}
