package hooks

import (
	"encoding/json"
	"fmt"
	. "github.com/xiabai84/githooks/config"
	. "github.com/xiabai84/githooks/types"
	. "github.com/xiabai84/githooks/utils"
	. "io/ioutil"
	"os"
	"strings"
	"text/template"
)

func AddGithooks(newHoook *GitHooks) {
	persistHooksAsLog(newHoook)
	createNewGitConfig(newHoook)
	updateGitConfigFile(newHoook)
}

func PreviewConfig(newHook *GitHooks) {
	previewGitConfigFile(newHook)
	previewNewGitConfig(newHook)
}

func previewGitConfigFile(hooks *GitHooks) {
	viewHeader := "========================== .gitconfig ==========================\n"
	bContent, err := ReadFile(GitConfigPath)
	if err != nil {
		fmt.Printf("Git configuration file %s doesn't exist, please setup this first.\n", GitConfigPath)
		os.Exit(-1)
	}
	configContent := string(bContent)
	tmpl, err := template.New("simple-hook-config").Funcs(template.FuncMap{
		"toLower": strings.ToLower,
	}).Parse(viewHeader + configContent + GitConfigPatch)
	CheckError(err)
	err = tmpl.Execute(os.Stdout, hooks)
	CheckError(err)
}

func previewNewGitConfig(hooks *GitHooks) {
	viewHeader := "========================== .gitconfig-" + strings.ToLower(hooks.Project) + " ==========================\n"
	tmpl, err := template.New("simple-jira-config").Parse(viewHeader + HooksConfigTmpl)
	CheckError(err)
	err = tmpl.Execute(os.Stdout, hooks)
	CheckError(err)
}

func persistHooksAsLog(hooks *GitHooks) {
	content, err := os.ReadFile(GithooksLogPath)
	CheckError(err)
	hookJson, _ := json.Marshal(hooks)
	newContent := string(content) + string(hookJson) + "\n"
	err = os.WriteFile(GithooksLogPath, []byte(newContent), 0755)
	CheckError(err)
}

func createNewGitConfig(hooks *GitHooks) {
	gitConfigPath := GitConfigPath + "-" + strings.ToLower(hooks.Project)
	tmpl, err := template.New("jira-config").Parse(HooksConfigTmpl)
	CheckError(err)
	f, err := os.Create(gitConfigPath)
	CheckError(err)
	err = tmpl.Execute(f, hooks)
	CheckError(err)
	err = f.Close()
	CheckError(err)
	fmt.Println("✅  Create new file:", gitConfigPath)
}

func updateGitConfigFile(hooks *GitHooks) {
	bContent, err := ReadFile(GitConfigPath)
	if err != nil {
		fmt.Printf("Git configuration file %s doesn't exist, please setup this first.\n", GitConfigPath)
		os.Exit(-1)
	}
	configContent := string(bContent)
	tmpl, err := template.New("simple-hook-config").Funcs(template.FuncMap{
		"toLower": strings.ToLower,
	}).Parse(configContent + GitConfigPatch)
	CheckError(err)
	f, err := os.Create(GitConfigPath)
	CheckError(err)
	err = tmpl.Execute(f, hooks)
	CheckError(err)
	err = f.Close()
	CheckError(err)
	fmt.Println("✅  Updated file:", GitConfigPath)
}
