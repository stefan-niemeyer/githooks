package types

import (
	"bufio"
	"encoding/json"
	"fmt"
	. "github.com/xiabai84/githooks/config"
	. "github.com/xiabai84/githooks/utils"
	. "io/ioutil"
	"os"
	"strings"
	"text/template"
)

type GitHooks struct {
	Project  string
	JiraName string
	WorkDir  string
}

// ConfigureHookFile create .gitconfig-<project> file under home dir
func (hooks *GitHooks) ConfigureHookFile() {
	tmpl, err := template.New(".gitconfig-project").Funcs(
		template.FuncMap{"upperCase": strings.ToUpper}).Parse(GithookTmpl)
	CheckError(err)
	homeDir, err := os.UserHomeDir()
	CheckError(err)
	filePath := homeDir + "/.gitconfig-" + strings.ToLower(hooks.Project)
	_, errorMsg := os.Stat(filePath)
	if errorMsg != nil {
		f, err := os.Create(filePath)
		CheckError(err)
		err = tmpl.Execute(f, &hooks)
		CheckError(err)
		err = f.Close()
		CheckError(err)
		fmt.Println("Created file ", filePath)
	} else {
		fmt.Printf("File %s already exists.", filePath)
	}

}

func (hooks *GitHooks) ConfigureGitConfig() {
	hasSlash := strings.HasSuffix(hooks.WorkDir, "/")
	if !hasSlash {
		hooks.WorkDir = hooks.WorkDir + "/"
	}
	homeDir, _ := os.UserHomeDir()
	configPath := homeDir + "/.gitconfig"
	_, errorMsg := os.Stat(configPath)
	if errorMsg == nil {
		hooks.UpdateCurrentGitConfig()
		fmt.Println("Updated .gitconfig.")
	}
}

func (hooks *GitHooks) ConfigureCommitMsg() {
	homeDir, _ := os.UserHomeDir()
	hookDir := homeDir + "/.githooks"
	commitMsgPath := hookDir + "/commit-msg"
	CreateDirIfNotExists(hookDir)
	_, errorMsg := os.Stat(commitMsgPath)
	if errorMsg != nil {
		tmpl, err := template.New(".githooks").Parse(CommitMsg)
		f, err := os.Create(commitMsgPath)
		CheckError(err)
		err = os.Chmod(commitMsgPath, 0755)
		CheckError(err)
		err = tmpl.Execute(f, &hooks)
		CheckError(err)
		err = f.Close()
		CheckError(err)
		fmt.Println("Created file ./githooks/commit-msg")
	}
}

func (hooks *GitHooks) DeleteHookGitConfig() {
	homeDir, _ := os.UserHomeDir()
	configPath := homeDir + "/.gitconfig-" + strings.ToLower(hooks.Project)
	err := os.Remove(configPath)
	CheckError(err)
}

func (hooks *GitHooks) PersistHooksAsLog() {
	gitHome := GetGithooksHome()
	hookLogPath := gitHome + "/" + GithooksLognName
	content, err := os.ReadFile(hookLogPath)
	CheckError(err)
	hookJson, _ := json.Marshal(hooks)
	newContent := string(content) + string(hookJson) + "\n"
	err = os.WriteFile(hookLogPath, []byte(newContent), 0755)
	CheckError(err)
}

func (hooks *GitHooks) UpdateCurrentGitConfig() {
	homeDir, err := os.UserHomeDir()
	gitConfigPath := homeDir + "/.gitconfig"
	CheckError(err)
	bContent, err := ReadFile(gitConfigPath)
	CheckError(err)
	configContent := string(bContent)
	tmpl, err := template.New("simple-hook-config").Funcs(template.FuncMap{
		"toUpper": strings.ToUpper,
		"toLower": strings.ToLower,
	}).Parse(configContent + GitConfigPatch)
	CheckError(err)
	f, err := os.Create(gitConfigPath)
	CheckError(err)
	err = tmpl.Execute(f, &hooks)
	CheckError(err)
	err = f.Close()
	CheckError(err)
	fmt.Println("✅  Updated file:", gitConfigPath)
}

func (hooks *GitHooks) RemoveCurrentHookFromLog() {
	gitHome := GetGithooksHome()
	hookLogPath := gitHome + "/" + GithooksLognName
	content, err := os.ReadFile(hookLogPath)
	CheckError(err)
	hooksArr := strings.Split(string(content), "\n")
	var removeTag int
	for idx, entry := range hooksArr {
		if entry == "" {
			continue
		}
		var delEntry = GitHooks{}
		err := json.Unmarshal([]byte(entry), &delEntry)
		CheckError(err)
		if delEntry.Project == hooks.Project {
			removeTag = idx
		}
	}
	newHooksArr := remove(hooksArr, removeTag)
	err = writeArrAsLines(newHooksArr, hookLogPath)
	CheckError(err)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func writeArrAsLines(lines []string, path string) error {
	file, err := os.Create(path)
	CheckError(err)
	defer func(file *os.File) {
		err := file.Close()
		CheckError(err)
	}(file)
	w := bufio.NewWriter(file)
	for _, line := range lines {
		if line != "" {
			_, err := fmt.Fprintln(w, line)
			CheckError(err)
		}
	}
	return w.Flush()
}
