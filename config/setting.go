package config

import (
	"os"
	"runtime"
)

const GitHooksFolder = ".githooks"
const GitHooksConfigFolder = "config"
const GithooksConfigName = "githooks.json"
const GitConfigFilename = ".gitconfig"
const GitHooksConfigPrefix = "gitconfig"
const AllowedTypesDefault = "build,chore,ci,docs,feat,fix,ops,perf,refactor,style,test"

var HomeDir, _ = os.UserHomeDir()
var HookDir = HomeDir + "/" + GitHooksFolder
var HookConfigDir = HookDir + "/" + GitHooksConfigFolder
var GithooksConfigFile = HookConfigDir + "/" + GithooksConfigName
var CommitMsgFile = HookDir + "/" + CommitMsgName()
var GitConfigFile = HomeDir + "/" + GitConfigFilename

var GitConfigPatch = `[includeIf "gitdir:{{ .Folder }}"]
    path = ` + GitHooksFolder + `/` + GitHooksConfigFolder + `/` + GitHooksConfigPrefix + `-{{ toLower .Name }}
`

var HooksConfigTmpl = `[core]
    hooksPath=~/` + GitHooksFolder + `
[user]
    jiraProjects={{ .ProjectKeyRE }}
    commitMessageStyle={{ .CommitMessageStyle }}
{{- if eq .CommitMessageStyle "conventional" }}
    allowedTypes={{ .AllowedTypes }}
{{- end }}
`

var DetailTmpl = `
{{ if ne .Name "Quit" }}
------------------ Workspace Configuration --------------------
Name: {{ .Name | faint }}
Folder: {{ .Folder | faint }}
Jira Project Key RegEx: {{ .ProjectKeyRE | faint }}
Commit Message Style: {{ .CommitMessageStyle | faint }}
{{- if eq .CommitMessageStyle "conventional" }}
Allowed AllowedTypes: {{ .AllowedTypes | faint }}
{{- end }}
{{ end }}
`

func CommitMsgName() string {
	if runtime.GOOS == "windows" {
		return "commit-msg.exe"
	}

	return "commit-msg"
}
