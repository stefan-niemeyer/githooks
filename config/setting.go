package config

const GithooksLognName = "githooks.log"

var GitConfigPatch = `[includeIf "gitdir:{{ .WorkDir }}"]
    path = .gitconfig-{{ toLower .Project }}
`

var HooksConfigTmpl = `
[core]
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
