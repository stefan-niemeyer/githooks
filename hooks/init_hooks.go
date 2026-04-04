package hooks

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/stefan-niemeyer/githooks/buildInfo"
	. "github.com/stefan-niemeyer/githooks/config"
	. "github.com/stefan-niemeyer/githooks/types"
	. "github.com/stefan-niemeyer/githooks/utils"
)

func InitHooks() GitHookConfig {
	ghConfig := GitHookConfig{
		Version:    buildInfo.GetBuildInfo().Version,
		Workspaces: []Workspace{},
	}
	CreateDirIfNotExists(HookDir)
	CreateDirIfNotExists(HookConfigDir)

	_, errorGitConfig := os.Stat(GitConfigFile)
	if errorGitConfig != nil {
		f, err := os.Create(GitConfigFile)
		CheckError(err)
		err = os.Chmod(GitConfigFile, 0644)
		CheckError(err)
		err = f.Close()
		CheckError(err)
		fmt.Println(promptui.IconGood+"  Created file", GitConfigFile)
	}

	// we copy the file again under Windows do assure that githooks and its copy have the same version
	abs, err := filepath.Abs(os.Args[0])
	CheckError(err)
	CloneOrLink(abs, CommitMsgFile, 0755)

	WriteGitHooksConfig(&ghConfig)
	fmt.Println(promptui.IconGood+"  Created file", GithooksConfigFile)

	return ghConfig
}
