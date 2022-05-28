package hooks

import (
	"encoding/json"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/stefan-niemeyer/githooks/buildInfo"
	. "github.com/stefan-niemeyer/githooks/config"
	. "github.com/stefan-niemeyer/githooks/types"
	. "github.com/stefan-niemeyer/githooks/utils"
	"io/ioutil"
	"os"
	"strings"
)

func MigrateGitHooksConfig() {
	_, errLogExists := os.Stat(GithooksLogPath)
	if errLogExists != nil {
		return
	}

	hookArr := ReadFromGitHooksLog()
	workspaces := make([]Workspace, 0, len(hookArr))
	for _, hook := range hookArr {
		workspaces = append(workspaces, Workspace{
			Name:         hook.Project,
			ProjectKeyRE: hook.JiraName,
			Folder:       hook.WorkDir,
		})
	}
	ghConfig := GitHookConfig{
		Version:    buildInfo.GetBuildInfo().Version,
		Workspaces: workspaces,
	}
	WriteGitHooksConfig(&ghConfig)
	fmt.Printf(promptui.IconGood+" Config file '%s' migrated to '%s'\n", GithooksLogPath, GithooksConfigPath)

	err := os.Remove(GithooksLogPath)
	CheckError(err)
	fmt.Printf(promptui.IconGood+" Old config file '%s' deleted\n", GithooksLogPath)
}

func WriteGitHooksConfig(ghConfig *GitHookConfig) {
	ghConfig.Version = buildInfo.GetBuildInfo().Version
	configJson, _ := json.Marshal(ghConfig)
	err := os.WriteFile(GithooksConfigPath, configJson, 0755)
	CheckError(err)
}

func ReadFromGitHooksLog() []GitHooks {
	var hookArr []GitHooks

	bytesRead, _ := ioutil.ReadFile(GithooksLogPath)
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		var hook = GitHooks{}
		err := json.Unmarshal([]byte(line), &hook)
		CheckError(err)
		hookArr = append(hookArr, hook)
	}

	return hookArr
}

func ReadGitHooksConfig() GitHookConfig {
	bytesRead, _ := ioutil.ReadFile(GithooksConfigPath)
	ghConfig := GitHookConfig{}
	err := json.Unmarshal(bytesRead, &ghConfig)
	CheckError(err)
	return ghConfig
}
