/*
   Created 2022 https://github.com/stefan-niemeyer
            and https://github.com/xiabai84
*/

package main

import (
	"os"
	"path/filepath"

	"github.com/stefan-niemeyer/githooks/cmd"
	"github.com/stefan-niemeyer/githooks/hooks"
	"github.com/stefan-niemeyer/githooks/utils"
)

func main() {
	abs, err := filepath.Abs(os.Args[0])
	utils.CheckError(err)
	basename := filepath.Base(abs)
	switch basename {
	case "commit-msg", "commit-msg.exe":
		hooks.CommitMsg(os.Args[1])
	default:
		cmd.Execute()
	}
}
