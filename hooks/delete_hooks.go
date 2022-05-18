package hooks

import (
	"bufio"
	"bytes"
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

func DeleteSelectedProject(hooks *GitHooks) {
	overwriteGitconfig(hooks)
	removeCurrentHookFromLog(hooks)
	deleteHookGitConfig(hooks)
	fmt.Println("✅  Removed project", hooks.Project)
}

func removeCurrentHookFromLog(hooks *GitHooks) {
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

func overwriteGitconfig(hooks *GitHooks) {
	homeDir, err := os.UserHomeDir()
	CheckError(err)
	gitConfigPath := homeDir + "/.gitconfig"
	bytesRead, _ := ReadFile(gitConfigPath)
	gitConfigContent := string(bytesRead)
	var partToReplace bytes.Buffer
	tmpl, err := template.New("origi").Funcs(template.FuncMap{
		"toLower": strings.ToLower,
	}).Parse(GitConfigPatch)
	CheckError(err)
	err = tmpl.Execute(&partToReplace, hooks)
	newGitConfigContent := strings.Replace(gitConfigContent, partToReplace.String(), "", -1)
	f, err := os.OpenFile(gitConfigPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	CheckError(err)
	_, err = f.Write([]byte(newGitConfigContent))
	CheckError(err)
	err = f.Close()
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

func deleteHookGitConfig(hooks *GitHooks) {
	homeDir, _ := os.UserHomeDir()
	configPath := homeDir + "/.gitconfig-" + strings.ToLower(hooks.Project)
	err := os.Remove(configPath)
	CheckError(err)
}
