package types

type GitHooks struct {
	Project  string
	JiraName string
	WorkDir  string
}

type GitHookConfig struct {
	Version    string      `json:"version,omitempty"`
	Workspaces []Workspace `json:"workspaces,omitempty"`
}

type Workspace struct {
	Name         string `json:"name,omitempty"`
	ProjectKeyRE string `json:"projectKeyRE,omitempty"`
	Folder       string `json:"folder,omitempty"`
}
