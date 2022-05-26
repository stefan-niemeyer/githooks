package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stefan-niemeyer/githooks/buildInfo"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Prints the githooks version",
	Long:    `Prints the githooks version`,
	Example: `githooks version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(buildInfo.GetBuildInfo().ToString())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringP("version", "v", "", "Prints the githooks version")
}
