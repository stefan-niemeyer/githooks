package cmd

import (
	"github.com/spf13/cobra"
	. "github.com/xiabai84/githooks/hooks"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all from githooks managed projects.",
	Long:  `List all githooks managed projects`,
	Run: func(cmd *cobra.Command, args []string) {
		ListAndSelectOne()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringP("list", "l", "", "List all githooks projects")
}
