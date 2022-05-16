package hooks

import (
	"encoding/json"
	"github.com/manifoldco/promptui"
	. "github.com/xiabai84/githooks/config"
	. "github.com/xiabai84/githooks/types"
	. "github.com/xiabai84/githooks/utils"
	"io/ioutil"
	"strings"
)

var detailTmpl = `
{{ if ne .Project "Quit" }}
------------------ Project's Configuration --------------------
Jira Project Name: {{ .Project | faint }}
Project Workspace: {{ .WorkDir | faint }}
{{ end }}
`

func ReadFromGitHookLog() []GitHooks {
	hookHome := GetGithooksHome()
	hookLog := hookHome + "/" + GithooksLognName
	var hookArr []GitHooks
	bytesRead, _ := ioutil.ReadFile(hookLog)
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

func ListAndSelectOne() GitHooks {
	hookArr := ReadFromGitHookLog()
	empty := GitHooks{Project: "Quit"}
	hookArr = append(hookArr, empty)

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "➣ {{ .Project | cyan }}",
		Inactive: "  {{ .Project | cyan }}",
		Selected: "➣ {{ .Project | red | cyan }}",
		Details:  detailTmpl,
	}

	searcher := func(input string, index int) bool {
		hook := hookArr[index]
		projName := strings.Replace(strings.ToLower(hook.Project), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(projName, input)
	}

	prompt := promptui.Select{
		Label:     "Active Githooks config file(s):",
		Items:     hookArr,
		Templates: templates,
		Size:      5,
		Searcher:  searcher,
	}
	i, _, err := prompt.Run()
	CheckError(err)
	return hookArr[i]
}
