package hooks

import (
	"bufio"
	"bytes"
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

func DeleteSelectedProject(hooks *GitHooks) {
	overwriteGitconfig(hooks)
	removeCurrentHookFromLog(hooks)
	deleteHookGitConfig(hooks)
	fmt.Println(promptui.IconGood+"  Removed project", hooks.Project)
}

func removeCurrentHookFromLog(hooks *GitHooks) {
	content, err := os.ReadFile(GithooksLogPath)
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
	err = writeArrAsLines(newHooksArr, GithooksLogPath)
	CheckError(err)
}

func overwriteGitconfig(hooks *GitHooks) {
	bytesRead, _ := ReadFile(GitConfigPath)
	gitConfigContent := string(bytesRead)
	var partToReplace bytes.Buffer
	tmpl, err := template.New("origi").Funcs(template.FuncMap{
		"toLower": strings.ToLower,
	}).Parse(GitConfigPatch)
	CheckError(err)
	err = tmpl.Execute(&partToReplace, hooks)
	newGitConfigContent := strings.Replace(gitConfigContent, partToReplace.String(), "", -1)
	f, err := os.OpenFile(GitConfigPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
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
	configPath := GitConfigPath + "-" + strings.ToLower(hooks.Project)
	err := os.Remove(configPath)
	CheckError(err)
}
