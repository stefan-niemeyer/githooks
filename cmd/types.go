package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var typesCmd = &cobra.Command{
	Use:     "types [type]",
	Aliases: []string{"show-types"},
	Short:   "Show supported Conventional Commit types",
	Long:    "Show supported Conventional Commit types and a short explanation for each one.",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		types := map[string]string{
			"build":    "Changes that affect the build system or external dependencies.",
			"chore":    "Routine maintenance tasks that do not affect application behavior.",
			"ci":       "Changes to continuous integration configuration or pipeline definitions.",
			"docs":     "Documentation-only changes.",
			"feat":     "Adds a new feature or user-facing functionality.",
			"fix":      "Fixes a bug or incorrect behavior.",
			"ops":      "Operational changes such as deployment, infrastructure, or monitoring updates.",
			"perf":     "Changes that improve performance.",
			"refactor": "Code changes that restructure implementation without changing behavior.",
			"revert":   "Reverts a previous commit or change.",
			"style":    "Formatting or style-only changes that do not affect logic.",
			"test":     "Adds or updates tests without changing production behavior.",
		}

		order := []string{"build", "chore", "ci", "docs", "feat", "fix", "ops", "perf", "refactor", "revert", "style", "test"}

		if len(args) == 1 {
			t, ok := types[args[0]]
			if !ok {
				return fmt.Errorf("unknown type: %s", args[0])
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%-10s %s\n", args[0], t)
			return nil
		}

		for _, name := range order {
			fmt.Fprintf(cmd.OutOrStdout(), "%-10s %s\n", name, types[name])
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(typesCmd)
}
