package hooks

import (
	"fmt"
	. "github.com/xiabai84/githooks/config"
	. "github.com/xiabai84/githooks/types"
	. "github.com/xiabai84/githooks/utils"
	"os"
)

func InitHooks() GitHooks {
	projName := GetPromptInput(Dialog{
		ErrorMsg: "❌ Please provide a Jira project key to track.",
		Label:    "🌟 Enter your project's name:",
	})

	workDir := GetPromptInput(Dialog{
		ErrorMsg: "❌ Please a path to workspace.",
		Label:    fmt.Sprintf("🌟 Enter your workspace:"),
	})

	userName := GetPromptInput(Dialog{
		ErrorMsg: "❌ Please provide your login name.",
		Label:    fmt.Sprintf("🌟 Enter git username:"),
	})

	email := GetPromptInput(Dialog{
		ErrorMsg: "❌ Please provide your login Email.",
		Label:    fmt.Sprintf("🌟 Enter your Email:"),
	})

	token := GetPromptInput(Dialog{
		ErrorMsg: "❌ Please provide a valid git token.",
		Label:    fmt.Sprintf("🌟 Enter git token:"),
	})

	hook := GitHooks{
		Project: projName,
		WorkDir: workDir,
		Name:    userName,
		Email:   email,
		Token:   token,
	}

	gitHome := GetGithooksHome()
	doesExist := CreateDirIfNotExists(gitHome)
	if !doesExist {
		_, err := os.Create(gitHome + "/" + GithooksLognName)
		CheckError(err)
	}
	hook.ConfigureHookFile()
	hook.ConfigureGitConfig()
	hook.ConfigureCommitMsg()
	hook.PersistHooksAsLog()
	return hook
}
