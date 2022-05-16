package hooks

import (
	"errors"
	. "github.com/xiabai84/githooks/config"
	. "github.com/xiabai84/githooks/types"
	. "github.com/xiabai84/githooks/utils"
	"os"
)

func InitHooks() GitHooks {
	hook := GitHooks{}
	gitHome := GetGithooksHome()
	doesExist := CreateDirIfNotExists(gitHome)
	if !doesExist {
		githooksLogPath := gitHome + "/" + GithooksLognName
		if _, err := os.Stat(githooksLogPath); errors.Is(err, os.ErrNotExist) {
			_, err := os.Create(githooksLogPath)
			CheckError(err)
		}
	}
	hook.ConfigureCommitMsg()
	return hook
}
