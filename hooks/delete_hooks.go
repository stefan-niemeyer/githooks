package hooks

import (
	"bytes"
	"fmt"
	"github.com/manifoldco/promptui"
	. "github.com/xiabai84/githooks/config"
	. "github.com/xiabai84/githooks/types"
	. "github.com/xiabai84/githooks/utils"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func DeleteSelectedProject(hook GitHooks) {
	homeDir, err := os.UserHomeDir()
	CheckError(err)
	gitConfigPath := homeDir + "/.gitconfig"
	bytesRead, _ := ioutil.ReadFile(gitConfigPath)
	gitConfigContent := string(bytesRead)
	var partToReplace bytes.Buffer
	tmpl, err := template.New("origi").Funcs(template.FuncMap{
		"toLower": strings.ToLower,
	}).Parse(GitConfigPatch)
	CheckError(err)
	err = tmpl.Execute(&partToReplace, hook)
	newGitConfigContent := strings.Replace(gitConfigContent, partToReplace.String(), "", -1)
	f, err := os.OpenFile(gitConfigPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	CheckError(err)
	_, err = f.Write([]byte(newGitConfigContent))
	CheckError(err)
	err = f.Close()
	CheckError(err)
	fmt.Println("✅  Removed project", hook.Project)
	hook.RemoveCurrentHookFromLog()
	hook.DeleteHookGitConfig()
}

func GetSelectedProject() GitHooks {
	hookArr := ReadFromGitHookLog()
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "➣ {{ .Project | cyan }}",
		Inactive: "  {{ .Project | cyan }}",
		Selected: "➣ {{ .Project | red | cyan }}",
		Details:  detailTmpl,
	}

	searcher := func(input string, index int) bool {
		hook := hookArr[index]
		name := strings.Replace(strings.ToLower(hook.Project), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}
	prompt1 := promptui.Select{
		Label:     "Delete:",
		Items:     hookArr,
		Templates: templates,
		Size:      5,
		Searcher:  searcher,
	}
	i, _, err := prompt1.Run()
	CheckError(err)

	prompt2 := promptui.Prompt{
		Label:     "Really want to delete this project",
		IsConfirm: true,
	}
	confirmed, err := prompt2.Run()
	if err != nil {
		fmt.Println("Canceled")
	}
	if confirmed != "y" {
		os.Exit(-1)
	}
	return hookArr[i]
}
