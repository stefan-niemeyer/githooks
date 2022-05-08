package utils

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func CheckError(e error) {
	if e != nil {
		_, err := fmt.Fprintln(os.Stderr, e)
		if err != nil {
			fmt.Println(e)
			os.Exit(1)
		}
		os.Exit(1)
	}
}

func CheckArgs(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		err := cmd.Help()
		CheckError(err)
		return
	}
}

func CreateDirIfNotExists(dirName string) bool {
	err := os.Mkdir(dirName, os.ModeDir)
	if err == nil {
		err := os.Chmod(dirName, 0755)
		CheckError(err)
		return false
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return false
		}
		if !info.IsDir() {
			return false
		}
		return true
	}
	return true
}

func GetGithooksHome() string {
	home, err := os.UserHomeDir()
	CheckError(err)
	return home + "/.githooks"
}
