package hooks

import (
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

func AddWorkspace(newWorkspace *Workspace) {
	persistConfigAsJson(newWorkspace)
	createWorkspaceGitConfig(newWorkspace)
	updateGitConfigFile(newWorkspace)
}

func PreviewConfig(newWorkspace *Workspace) {
	previewGitConfigFile(newWorkspace)
	previewWorkspaceGitConfig(newWorkspace)
}

func CheckConfigFiles() {
	_, err1 := os.Stat(GitConfigPath)
	_, err2 := os.Stat(HookDir)
	_, err3 := os.Stat(HookConfigDir)
	_, err4 := os.Stat(CommitMsgPath)
	_, err5 := os.Stat(GithooksConfigPath)

	switch {
	case err1 != nil:
		fmt.Printf(promptui.IconBad+" File %s doesn't exist, please execute 'githooks init' first.\n", GitConfigPath)
		os.Exit(1)

	case err2 != nil:
		fmt.Printf(promptui.IconBad+" File %s doesn't exist, please execute 'githooks init' first.\n", HookDir)
		os.Exit(1)

	case err3 != nil:
		fmt.Printf(promptui.IconBad+" File %s doesn't exist, please execute 'githooks init' first.\n", HookConfigDir)
		os.Exit(1)

	case err4 != nil:
		fmt.Printf(promptui.IconBad+" File %s doesn't exist, please execute 'githooks init' first.\n", CommitMsgPath)
		os.Exit(1)

	case err5 != nil:
		fmt.Printf(promptui.IconBad+" File %s doesn't exist, please execute 'githooks init' first.\n", GithooksConfigPath)
		os.Exit(1)
	}
}

func previewGitConfigFile(workspace *Workspace) {
	viewHeader := "========================== ~/" + GitConfigFilename + " ==========================\n"
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
	err = tmpl.Execute(os.Stdout, workspace)
	CheckError(err)
}

func previewWorkspaceGitConfig(workspace *Workspace) {
	viewHeader := "========================== ~/" + GitHooksFolder + "/" + GitHooksConfigFolder + "/" + GitHooksConfigPraefix + "-" + strings.ToLower(workspace.Name) + " ==========================\n"
	tmpl, err := template.New("simple-jira-config").Parse(viewHeader + HooksConfigTmpl)
	CheckError(err)
	err = tmpl.Execute(os.Stdout, workspace)
	CheckError(err)
}

func persistConfigAsJson(workspace *Workspace) {
	ghConfig := ReadGitHooksConfig()
	ghConfig.Workspaces = append(ghConfig.Workspaces, *workspace)
	WriteGitHooksConfig(&ghConfig)
}

func createWorkspaceGitConfig(workspace *Workspace) {
	workspaceGitConfigPath := HookConfigDir + "/" + GitHooksConfigPraefix + "-" + strings.ToLower(workspace.Name)
	tmpl, err := template.New("jira-config").Parse(HooksConfigTmpl)
	CheckError(err)
	f, err := os.Create(workspaceGitConfigPath)
	CheckError(err)
	err = tmpl.Execute(f, workspace)
	CheckError(err)
	err = f.Close()
	CheckError(err)
	fmt.Println(promptui.IconGood+"  Create new file:", workspaceGitConfigPath)
}

func updateGitConfigFile(workspace *Workspace) {
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
	err = tmpl.Execute(f, workspace)
	CheckError(err)
	err = f.Close()
	CheckError(err)
	fmt.Println(promptui.IconGood+"  Updated file:", GitConfigPath)
}
