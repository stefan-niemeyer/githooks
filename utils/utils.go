package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"runtime"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
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

func CloneOrLink(src, dst string, mode fs.FileMode) {
	if runtime.GOOS == "windows" {
		// real copy
		srcExe := addExeIfNeeded(src)
		data, err := os.ReadFile(srcExe)
		CheckError(err)
		dstExe := addExeIfNeeded(dst)
		err = os.WriteFile(dstExe, data, mode)
		CheckError(err)
		fmt.Println(promptui.IconGood+"  Created copy", dstExe)
	} else {
		err := os.Remove(dst)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			CheckError(err)
		}

		// symlink on not-Windows
		err = os.Symlink(src, dst)
		CheckError(err)
		err = os.Chmod(dst, mode)
		CheckError(err)
		fmt.Println(promptui.IconGood+"  Created link", dst)
	}
}

func addExeIfNeeded(s string) string {
	exeLower := ".exe"
	if strings.HasSuffix(strings.ToLower(s), exeLower) {
		return s
	}
	return s + exeLower
}
