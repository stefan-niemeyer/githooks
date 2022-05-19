package hooks

import (
	"fmt"
	. "github.com/xiabai84/githooks/config"
	. "github.com/xiabai84/githooks/types"
	. "github.com/xiabai84/githooks/utils"
	"os"
	"text/template"
)

func InitHooks() GitHooks {
	hook := GitHooks{}
	CreateDirIfNotExists(HookDir)

	_, errorMsg := os.Stat(CommitMsgPath)
	if errorMsg != nil {
		tmpl, err := template.New(".githooks").Parse(CommitMsg)
		f, err := os.Create(CommitMsgPath)
		CheckError(err)
		err = os.Chmod(CommitMsgPath, 0755)
		CheckError(err)
		err = tmpl.Execute(f, &hook)
		CheckError(err)
		err = f.Close()
		CheckError(err)
		fmt.Println("✅  Created file", CommitMsgPath)
	}

	_, errorLog := os.Stat(GithooksLogPath)
	if errorLog != nil {
		_, err := os.Create(GithooksLogPath)
		CheckError(err)
		fmt.Println("✅  Created file" + GithooksLogPath)
	}

	return hook
}
