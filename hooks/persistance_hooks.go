package hooks

import (
	"encoding/json"
	"os"

	"github.com/stefan-niemeyer/githooks/buildInfo"
	. "github.com/stefan-niemeyer/githooks/config"
	"github.com/stefan-niemeyer/githooks/styles"
	. "github.com/stefan-niemeyer/githooks/types"
	. "github.com/stefan-niemeyer/githooks/utils"
)

func WriteGitHooksConfig(ghConfig *GitHookConfig) {
	ghConfig.Version = buildInfo.GetBuildInfo().Version
	configJson, _ := json.Marshal(ghConfig)
	err := os.WriteFile(GithooksConfigFile, configJson, 644)
	CheckError(err)
}

func ReadGitHooksConfig() GitHookConfig {
	bytesRead, _ := os.ReadFile(GithooksConfigFile)
	ghConfig := GitHookConfig{}
	err := json.Unmarshal(bytesRead, &ghConfig)
	for idx, workspace := range ghConfig.Workspaces {
		if len(workspace.CommitMessageStyle) == 0 {
			ghConfig.Workspaces[idx].CommitMessageStyle = string(styles.PlainBrackets)
		}
	}

	CheckError(err)
	return ghConfig
}
