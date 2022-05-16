package cmd

import (
	. "github.com/xiabai84/githooks/hooks"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a githooks project and its settings",
	Long:  `A longer description `,
	Run: func(cmd *cobra.Command, args []string) {
		DeleteSelectedProject()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
