package config

import (
	"os"
)

var HomeDir, _ = os.UserHomeDir()
var HookDir = HomeDir + "/.githooks"

const GithooksLogName = "githooks.log"
const GithooksConfigName = "githooks.json"
const CommitMsgName = "commit-msg"

var GithooksLogPath = HookDir + "/" + GithooksLogName
var GithooksConfigPath = HookDir + "/" + GithooksConfigName
var CommitMsgPath = HookDir + "/" + CommitMsgName
var GitConfigPath = HomeDir + "/.gitconfig"

var GitConfigPatch = `[includeIf "gitdir:{{ .Folder }}"]
    path = .gitconfig-{{ toLower .Name }}
`

var HooksConfigTmpl = `[core]
    hooksPath=~/.githooks
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
