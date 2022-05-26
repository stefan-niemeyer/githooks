package config

import (
	"os"
)

var HomeDir, _ = os.UserHomeDir()
var HookDir = HomeDir + "/.githooks"

const GithooksLogName = "githooks.log"
const CommitMsgName = "commit-msg"

var GithooksLogPath = HookDir + "/" + GithooksLogName
var CommitMsgPath = HookDir + "/" + CommitMsgName
var GitConfigPath = HomeDir + "/.gitconfig"

var GitConfigPatch = `[includeIf "gitdir:{{ .WorkDir }}"]
    path = .gitconfig-{{ toLower .Project }}
`

var HooksConfigTmpl = `[core]
    hooksPath=~/.githooks
[user]
    jiraProjects={{ .JiraName }}
`

var DetailTmpl = `
{{ if ne .Project "Quit" }}
------------------ Project's Configuration --------------------
Jira Project Name: {{ .Project | faint }}
Project Workspace: {{ .WorkDir | faint }}
{{ end }}
`
