package cmd

import (
	"fmt"
	"github.com/stefan-niemeyer/githooks/utils"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "githooks",
	Short: "githooks help developer on setting name's conventions of git-commit-msg ",
	Long:  `githooks prevent developer enter unexpected commit messages, which don't contain predefined name's conventions.'`,
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

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
