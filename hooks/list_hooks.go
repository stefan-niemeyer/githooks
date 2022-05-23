package hooks

import (
	"encoding/json"
	. "github.com/stefan-niemeyer/githooks/config"
	. "github.com/stefan-niemeyer/githooks/types"
	. "github.com/stefan-niemeyer/githooks/utils"
	"io/ioutil"
	"strings"
)

func ReadFromGitHookLog() []GitHooks {
	var hookArr []GitHooks

	bytesRead, _ := ioutil.ReadFile(GithooksLogPath)
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		var hook = GitHooks{}
		err := json.Unmarshal([]byte(line), &hook)
		CheckError(err)
		hookArr = append(hookArr, hook)
	}

	return hookArr
}
