package config

const GithooksLognName = "githooks.log"

var GithookTmpl = `
[core]
    hooksPath=~/.githooks
[user]
    jiraProjects={{ upperCase .Project }}
`

var GitConfigTmpl = `
[user]
	token = {{.Token}}
	name = {{.Name}}
	email = {{.Email}}

[includeIf "gitdir:{{.WorkDir}}"]
    path = .gitconfig-{{ lowerCase .Project }}

`

var GitConfigPatch = `[includeIf "gitdir:{{ .WorkDir }}"]
    path = .gitconfig-{{ toLower .Project }}
`
