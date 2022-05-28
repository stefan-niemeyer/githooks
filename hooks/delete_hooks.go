package hooks

import (
	"bytes"
	"fmt"
	"github.com/manifoldco/promptui"
	. "github.com/stefan-niemeyer/githooks/config"
	. "github.com/stefan-niemeyer/githooks/types"
	. "github.com/stefan-niemeyer/githooks/utils"
	. "io/ioutil"
	"os"
	"strings"
	"text/template"
)

func DeleteSelectedWorkspace(ghConfig *GitHookConfig, idx int) {
	removedWorkspace := ghConfig.Workspaces[idx].Name
	overwriteGitConfig(&ghConfig.Workspaces[idx])
	ghConfig.Workspaces = append(ghConfig.Workspaces[:idx], ghConfig.Workspaces[idx+1:]...)
	WriteGitHooksConfig(ghConfig)
	deleteWorkspaceGitConfig(removedWorkspace)
	fmt.Println(promptui.IconGood+"  Removed workspace", removedWorkspace)
}

func overwriteGitConfig(workspace *Workspace) {
	bytesRead, _ := ReadFile(GitConfigPath)
	gitConfigContent := string(bytesRead)
	var partToReplace bytes.Buffer
	tmpl, err := template.New("original").Funcs(template.FuncMap{
		"toLower": strings.ToLower,
	}).Parse(GitConfigPatch)
	CheckError(err)
	err = tmpl.Execute(&partToReplace, workspace)
	newGitConfigContent := strings.Replace(gitConfigContent, partToReplace.String(), "", -1)
	f, err := os.OpenFile(GitConfigPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	CheckError(err)
	_, err = f.Write([]byte(newGitConfigContent))
	CheckError(err)
	err = f.Close()
	CheckError(err)
}

func deleteWorkspaceGitConfig(wsName string) {
	configPath := GitConfigPath + "-" + strings.ToLower(wsName)
	err := os.Remove(configPath)
	CheckError(err)
}
