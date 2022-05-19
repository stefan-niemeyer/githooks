package hooks

import (
	"errors"
	"fmt"
	. "github.com/xiabai84/githooks/config"
	. "github.com/xiabai84/githooks/types"
	. "github.com/xiabai84/githooks/utils"
	"os"
	"text/template"
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

	homeDir, _ := os.UserHomeDir()
	hookDir := homeDir + "/.githooks"
	commitMsgPath := hookDir + "/commit-msg"
	githooksLogPath := hookDir + "/" + GithooksLognName
	CreateDirIfNotExists(hookDir)

	_, errorMsg := os.Stat(commitMsgPath)
	if errorMsg != nil {
		tmpl, err := template.New(".githooks").Parse(CommitMsg)
		f, err := os.Create(commitMsgPath)
		CheckError(err)
		err = os.Chmod(commitMsgPath, 0755)
		CheckError(err)
		err = tmpl.Execute(f, &hook)
		CheckError(err)
		err = f.Close()
		CheckError(err)
		fmt.Println("✅  Created file ./githooks/commit-msg")
	}

	_, errorLog := os.Stat(githooksLogPath)
	if errorLog != nil {
		_, err := os.Create(githooksLogPath)
		CheckError(err)
		fmt.Println("✅  Created file ./githooks/" + GithooksLognName)
	}

	return hook
}
