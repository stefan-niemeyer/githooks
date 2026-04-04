package prompt

import (
	"strings"

	. "github.com/stefan-niemeyer/githooks/types"
)

// newWorkspaceSearcher erstellt eine Suchfunktion für promptui.Select
func NewWorkspaceSearcher(workspaces []Workspace) func(string, int) bool {
	// Gibt die eigentliche Funktion mit der von promptui erwarteten Signatur zurück
	return func(input string, index int) bool {
		workspace := workspaces[index]

		// Tipp: strings.ReplaceAll ist in modernem Go performanter und
		// lesbarer als strings.Replace(..., -1)
		name := strings.ReplaceAll(strings.ToLower(workspace.Name), " ", "")
		input = strings.ReplaceAll(strings.ToLower(input), " ", "")

		return strings.Contains(name, input)
	}
}
