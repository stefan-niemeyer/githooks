package hooks

import (
	"os"
	"path/filepath"
	"runtime"

	. "github.com/stefan-niemeyer/githooks/config"
	. "github.com/stefan-niemeyer/githooks/utils"
)

func UpdateHooks() {
	CreateDirIfNotExists(HookDir)

	_, errorMsg := os.Stat(CommitMsgFile)
	if errorMsg != nil || runtime.GOOS == "windows" {
		// we copy the file again under Windows do assure that githooks and its copy have the same version
		abs, err := filepath.Abs(os.Args[0])
		CheckError(err)
		CloneOrLink(abs, CommitMsgFile, 0755)
	}
}
