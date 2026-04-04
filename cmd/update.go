package cmd

import (
	"github.com/spf13/cobra"
	. "github.com/stefan-niemeyer/githooks/hooks"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the githooks",
	Long:  `Update the githooks`,
	Run: func(cmd *cobra.Command, args []string) {
		UpdateHooks()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
