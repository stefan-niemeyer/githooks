package hooks

import (
	"encoding/json"
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

func AddGithooks(newHoook *GitHooks) {
	persistHooksAsLog(newHoook)
	createNewGitConfig(newHoook)
	updateGitConfigFile(newHoook)
}

func PreviewConfig(newHook *GitHooks) {
	previewGitConfigFile(newHook)
	previewNewGitConfig(newHook)
}

func CheckConfigFiles() {
	_, err1 := os.Stat(GitConfigPath)
	_, err2 := os.Stat(HookDir)
	_, err3 := os.Stat(CommitMsgPath)
	_, err4 := os.Stat(GithooksLogPath)

	switch {
	case err1 != nil:
		fmt.Printf(promptui.IconBad+" File %s doesn't exist, please execute 'githooks init' first.\n", GitConfigPath)
		os.Exit(1)

	case err2 != nil:
		fmt.Printf(promptui.IconBad+" File %s doesn't exist, please execute 'githooks init' first.\n", HookDir)
		os.Exit(1)

	case err3 != nil:
		fmt.Printf(promptui.IconBad+" File %s doesn't exist, please execute 'githooks init' first.\n", CommitMsgPath)
		os.Exit(1)

	case err4 != nil:
		fmt.Printf(promptui.IconBad+" File %s doesn't exist, please execute 'githooks init' first.\n", GithooksLogPath)
		os.Exit(1)
	}
}

func previewGitConfigFile(hooks *GitHooks) {
	viewHeader := "========================== .gitconfig ==========================\n"
	bContent, err := ReadFile(GitConfigPath)
	if err != nil {
		fmt.Printf("Git configuration file %s doesn't exist, please setup git first.\n", GitConfigPath)
		os.Exit(1)
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
	fmt.Println(promptui.IconGood+"  Create new file:", gitConfigPath)
}

func updateGitConfigFile(hooks *GitHooks) {
	bContent, err := ReadFile(GitConfigPath)
	if err != nil {
		fmt.Printf("Git configuration file %s doesn't exist, please setup this first.\n", GitConfigPath)
		os.Exit(1)
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
	fmt.Println(promptui.IconGood+"  Updated file:", GitConfigPath)
}
