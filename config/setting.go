package config

import (
	"os"
)

const GitHooksFolder = ".githooks"
const GitHooksConfigFolder = "config"
const GithooksLogName = "githooks.log"
const GithooksConfigName = "githooks.json"
const CommitMsgName = "commit-msg"
const GitConfigFilename = ".gitconfig"
const GitHooksConfigPraefix = "gitconfig"

var HomeDir, _ = os.UserHomeDir()
var HookDir = HomeDir + "/" + GitHooksFolder
var HookConfigDir = HookDir + "/" + GitHooksConfigFolder
var GithooksLogPath = HookConfigDir + "/" + GithooksLogName
var GithooksConfigPath = HookConfigDir + "/" + GithooksConfigName
var CommitMsgPath = HookDir + "/" + CommitMsgName
var GitConfigPath = HomeDir + "/" + GitConfigFilename

var GitConfigPatch = `[includeIf "gitdir:{{ .Folder }}"]
    path = ` + GitHooksFolder + `/` + GitHooksConfigFolder + `/` + GitHooksConfigPraefix + `-{{ toLower .Name }}
`

var HooksConfigTmpl = `[core]
    hooksPath=~/` + GitHooksFolder + `
[user]
    jiraProjects={{ .ProjectKeyRE }}
`

var DetailTmpl = `
{{ if ne .Name "Quit" }}
------------------ Workspace Configuration --------------------
Name: {{ .Name | faint }}
Folder: {{ .Folder | faint }}
Jira Project Key RegEx: {{ .ProjectKeyRE | faint }}
{{ end }}
`
