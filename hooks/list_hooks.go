package hooks

import (
	. "github.com/stefan-niemeyer/githooks/types"
	. "github.com/stefan-niemeyer/githooks/utils"
	"os"
	"strings"
)

func GetWorkspaceIndex(workspaces []Workspace) int {
	var matchIdx int

	home, errHome := os.UserHomeDir()
	CheckError(errHome)

	cwd, errCwd := os.Getwd()
	CheckError(errCwd)
	if len(cwd) == 0 {
		return matchIdx
	}
	cwd += "/"

	longestMatch := 0
	for idx, workspace := range workspaces {
		wsFolder := strings.Replace(workspace.Folder, "~", home, 1)
		if strings.HasPrefix(cwd, wsFolder) && len(wsFolder) > longestMatch {
			longestMatch = len(wsFolder)
			matchIdx = idx
		}
	}

	return matchIdx
}
