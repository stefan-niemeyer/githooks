package hooks

import (
	"fmt"
	"github.com/manifoldco/promptui"
	. "github.com/xiabai84/githooks/config"
	. "github.com/xiabai84/githooks/types"
	. "github.com/xiabai84/githooks/utils"
	. "io/ioutil"
	"os"
	"strings"
	"text/template"
)

func AddGithooks() {
	githooksHome := GetGithooksHome()
	hookLogPath := githooksHome + "/" + GithooksLognName
	_, err := os.ReadFile(hookLogPath)
	if err != nil {
		fmt.Printf("❌  File %s doesn't exist. Please perform githooks init first.", hookLogPath)
		os.Exit(-1)
	}

	projName := GetPromptInput(Dialog{
		ErrorMsg: "❌ Please provide a Jira project key to track.",
		Label:    "🌟 Enter your Jira project's name:",
	})

	workDir := GetPromptInput(Dialog{
		ErrorMsg: "❌ Please a path to workspace.",
		Label:    fmt.Sprintf("🌟 Enter path to your workspace:"),
	})
	hasSlash := strings.HasSuffix(workDir, "/")
	if !hasSlash {
		workDir = workDir + "/"
	}
	newHook := GitHooks{Project: projName, JiraName: strings.ToUpper(projName), WorkDir: workDir}
	PreviewGitConfigFile(newHook)
	PreviewNewGitConfig(newHook)

	prompt := promptui.Prompt{
		Label:     "Input was correct",
		IsConfirm: true,
	}

	confirmed, err := prompt.Run()
	if err != nil {
		fmt.Println("❌ Canceled setting new githooks project.")
	}
	if confirmed == "y" {
		newHook.PersistHooksAsLog()
		CreateNewGitConfig(newHook)
		UpdateGitConfigFile(newHook)
	}
}

func PreviewGitConfigFile(newHook GitHooks) {
	viewHeader := "========================== .gitconfig ==========================\n"
	homeDir, err := os.UserHomeDir()
	CheckError(err)
	bContent, err := ReadFile(homeDir + "/.gitconfig")
	CheckError(err)
	configContent := string(bContent)
	tmpl, err := template.New("simple-hook-config").Funcs(template.FuncMap{
		"toLower": strings.ToLower,
	}).Parse(viewHeader + configContent + GitConfigPatch)
	CheckError(err)
	err = tmpl.Execute(os.Stdout, newHook)
	CheckError(err)
}

func UpdateGitConfigFile(newHook GitHooks) {
	homeDir, err := os.UserHomeDir()
	gitConfigPath := homeDir + "/.gitconfig"
	CheckError(err)
	bContent, err := ReadFile(gitConfigPath)
	CheckError(err)
	configContent := string(bContent)
	tmpl, err := template.New("simple-hook-config").Funcs(template.FuncMap{
		"toLower": strings.ToLower,
	}).Parse(configContent + GitConfigPatch)
	CheckError(err)
	f, err := os.Create(gitConfigPath)
	CheckError(err)
	err = tmpl.Execute(f, newHook)
	CheckError(err)
	err = f.Close()
	CheckError(err)
	fmt.Println("✅  Updated file:", gitConfigPath)
}

func CreateNewGitConfig(newHook GitHooks) {
	homeDir, err := os.UserHomeDir()
	gitConfigPath := homeDir + "/.gitconfig-" + strings.ToLower(newHook.Project)
	tmpl, err := template.New("jira-config").Parse(ConfigJiraTmpl)
	CheckError(err)
	f, err := os.Create(gitConfigPath)
	CheckError(err)
	err = tmpl.Execute(f, newHook)
	CheckError(err)
	err = f.Close()
	CheckError(err)
	fmt.Println("✅  Create new file:", gitConfigPath)
}

func PreviewNewGitConfig(newHook GitHooks) {
	viewHeader := "========================== .gitconfig-" + strings.ToLower(newHook.Project) + " ==========================\n"
	tmpl, err := template.New("simple-jira-config").Parse(viewHeader + ConfigJiraTmpl)
	CheckError(err)
	err = tmpl.Execute(os.Stdout, newHook)
	CheckError(err)
}
