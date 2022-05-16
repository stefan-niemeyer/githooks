package hooks

import (
	"fmt"
	"github.com/manifoldco/promptui"
	. "github.com/xiabai84/githooks/types"
	. "github.com/xiabai84/githooks/utils"
	"os"
	"strings"
)

func DeleteSelectedProject() {
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

	if hookArr[i].Project != "Quit" {
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
		hookArr[i].OverwriteGitconfig()
		hookArr[i].RemoveCurrentHookFromLog()
		hookArr[i].DeleteHookGitConfig()
		fmt.Println("✅  Removed project", hookArr[i].Project)
	} else {
		os.Exit(0)
	}
}
