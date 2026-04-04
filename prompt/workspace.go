package prompt

import (
	"fmt"
	"os"
	"strings"

	"github.com/stefan-niemeyer/githooks/config"
	. "github.com/stefan-niemeyer/githooks/styles"
	. "github.com/stefan-niemeyer/githooks/types"
)

func CreateWorkspace(nameDefault, projectKeyREDefault, folderDefault, typesDefault string, formatStyleDefault FormatStyle) Workspace {
	label := "Enter your workspace name:"
	if len(nameDefault) != 0 {
		label = fmt.Sprintf("Enter your workspace name (%s):", nameDefault)
	}
	projName := GetPromptInput(Dialog{
		ErrorMsg: "Please provide a name for the workspace.",
		Label:    label,
	}, nameDefault, false)

	projectKeyRE := ""
	if len(projectKeyREDefault) != 0 {
		projectKeyRE = projectKeyREDefault
	}
	projectKeyRE = strings.ToUpper(GetPromptInput(Dialog{
		ErrorMsg: "Please provide a Jira project key RegEx to track, e.g. ALPHA or (ALPHA|BETA)",
		Label:    fmt.Sprintf("Enter your Jira project key RegEx, keep empty to accept any project key (%s):", projectKeyRE),
	}, projectKeyRE, true))

	cmStyle := SelectFormatStyle(formatStyleDefault)
	allowedTypes := config.AllowedTypesDefault
	if cmStyle == Conventional {
		if len(typesDefault) != 0 {
			allowedTypes = typesDefault
		}
		allowedTypes = GetPromptInput(Dialog{
			ErrorMsg: "Please provide a list of allowed allowedTypes, e.g. 'fix,feat,...'",
			Label:    fmt.Sprintf("Enter the allowed allowedTypes (%s):", allowedTypes),
		}, allowedTypes, false)
	}

	folderDefaultClean := folderDefault
	homeDir, err := os.UserHomeDir()
	if err == nil && len(homeDir) != 0 {
		folderDefaultClean = strings.Replace(folderDefault, homeDir, "~", 1)
	}
	folder := GetPromptInput(Dialog{
		ErrorMsg: "Please enter a path to your workspace.",
		Label:    fmt.Sprintf("Enter path to your workspace (%s):", folderDefault),
	}, folderDefaultClean, false)
	if len(homeDir) != 0 {
		folder = strings.Replace(folder, homeDir, "~", 1)
	}
	if !strings.HasSuffix(folder, "/") {
		folder += "/"
	}

	newWorkspace := Workspace{
		Name:               projName,
		ProjectKeyRE:       strings.ToUpper(projectKeyRE),
		CommitMessageStyle: string(cmStyle),
		Folder:             folder,
		AllowedTypes:       allowedTypes,
	}

	return newWorkspace
}
