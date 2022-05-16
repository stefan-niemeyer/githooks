package config

const GithooksLognName = "githooks.log"

var GithookTmpl = `
[core]
    hooksPath=~/.githooks
[user]
    jiraProjects={{ upperCase .Project }}
`

var GitConfigPatch = `[includeIf "gitdir:{{ .WorkDir }}"]
    path = .gitconfig-{{ toLower .Project }}
`

var ConfigJiraTmpl = `
[core]
    hooksPath=~/.githooks
[user]
    jiraProjects={{ .JiraName }}
`
