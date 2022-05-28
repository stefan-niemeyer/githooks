package hooks

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/stefan-niemeyer/githooks/buildInfo"
	. "github.com/stefan-niemeyer/githooks/config"
	. "github.com/stefan-niemeyer/githooks/types"
	. "github.com/stefan-niemeyer/githooks/utils"
	"os"
	"text/template"
)

func InitHooks() GitHookConfig {
	ghConfig := GitHookConfig{
		Version:    buildInfo.GetBuildInfo().Version,
		Workspaces: []Workspace{},
	}
	CreateDirIfNotExists(HookDir)

	_, errorGitConfig := os.Stat(GitConfigPath)
	if errorGitConfig != nil {
		f, err := os.Create(GitConfigPath)
		CheckError(err)
		err = os.Chmod(GitConfigPath, 0644)
		CheckError(err)
		err = f.Close()
		CheckError(err)
		fmt.Println(promptui.IconGood+"  Created file", GitConfigPath)
	}

	_, errorMsg := os.Stat(CommitMsgPath)
	if errorMsg != nil {
		tmpl, err := template.New(".githooks").Parse(CommitMsg)
		f, err := os.Create(CommitMsgPath)
		CheckError(err)
		err = os.Chmod(CommitMsgPath, 0755)
		CheckError(err)
		err = tmpl.Execute(f, &ghConfig)
		CheckError(err)
		err = f.Close()
		CheckError(err)
		fmt.Println(promptui.IconGood+"  Created file", CommitMsgPath)
	}

	WriteGitHooksConfig(&ghConfig)
	fmt.Println(promptui.IconGood+"  Created file", GithooksConfigPath)

	return ghConfig
}
