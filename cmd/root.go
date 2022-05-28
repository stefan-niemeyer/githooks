package cmd

import (
	"fmt"
	"github.com/stefan-niemeyer/githooks/hooks"
	"github.com/stefan-niemeyer/githooks/utils"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "githooks",
	Short: "githooks helps developers with setting name conventions of a Git commit message",
	Long:  `githooks prevents developer to enter commit messages, which don't contain predefined Jira issue keys.'`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.CheckArgs(cmd, args)
	},
}

func Execute() {
	err := rootCmd.Execute()
	utils.CheckError(err)
}

func init() {
	cobra.OnInitialize(initConfig)
	hooks.MigrateGitHooksConfig()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".githooks" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".githooks")
	}

	viper.AutomaticEnv() // read environment variables that match

	// If a config file is found, read it.
	if err := viper.ReadInConfig(); err == nil {
		_, err := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		utils.CheckError(err)
	}
}
