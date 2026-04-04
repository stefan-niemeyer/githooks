package types

type GitHooks struct {
	Project  string
	JiraName string
	CMStyle  string
	WorkDir  string
}

type GitHookConfig struct {
	Version    string      `json:"version,omitempty"`
	Workspaces []Workspace `json:"workspaces,omitempty"`
}

type Workspace struct {
	Name               string `json:"name,omitempty"`
	ProjectKeyRE       string `json:"projectKeyRE,omitempty"`
	CommitMessageStyle string `json:"commitMessageStyle,omitempty"`
	AllowedTypes       string `json:"allowedtypes,omitempty"`
	Folder             string `json:"folder,omitempty"`
}
