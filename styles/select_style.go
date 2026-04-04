package styles

import (
	"os"
	"slices"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/stefan-niemeyer/githooks/config"
	"github.com/stefan-niemeyer/githooks/utils"
)

// Any type can be given to the select's item as long as the templates properly implement the dot notation
// to display it.
type cmStyle struct {
	Name               string
	CommitMessageStyle FormatStyle
	MessageBefore      string
	MessageAfter       string
}

func FindStyleIndex(styles []cmStyle, targetStyle FormatStyle) int {
	// slices.IndexFunc returns -1 if nothing was found
	return slices.IndexFunc(styles, func(s cmStyle) bool {
		return s.CommitMessageStyle == targetStyle
	})
}

func SelectFormatStyle(styleDefault FormatStyle) FormatStyle {
	// The select will show a series of peppers stored inside a slice of structs. To display the content of the struct,
	// the usual dot notation is used inside the templates to select the fields and color them.
	cmStyles := []cmStyle{
		{"Plain Space", PlainSpace, "add feature", "JIRA-123 add feature"},
		{"Plain Colon", PlainColon, "add feature", "JIRA-123: add feature"},
		{"Plain Brackets", PlainBrackets, "add feature", "[JIRA-123] add feature"},
		{"Conventional", Conventional, "fix(parser)!: changes API", "fix(parser)!: changes API\n\nRefs: JIRA-123"},
		{"Quit", PlainSpace, "", ""},
	}

	styleIdx := FindStyleIndex(cmStyles, styleDefault)
	if styleIdx == -1 {
		styleIdx = FindStyleIndex(cmStyles, PlainBrackets)
	}
	// The Active and Selected templates set a small cmStyles icon next to the name colored and the heat unit for the
	// active template. The details template is show at the bottom of the select's list and displays the full info
	// for that cmStyles in a multi-line template.
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "➣ {{ .Name | cyan}}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "➣ {{ .Name | red | cyan }}",
		Details: `
{{ if ne .Name "Quit" }}
--------- Format Style ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Message Before:" | faint }}	{{ .MessageBefore }}
{{ "Message After: " | faint }}	{{ .MessageAfter }}
{{ end }}`,
	}

	// A searcher function is implemented which enabled the search mode for the select. The function follows
	// the required searcher signature and finds any cmStyles whose name contains the searched string.
	searcher := func(input string, index int) bool {
		style := cmStyles[index]
		name := strings.Replace(strings.ToLower(style.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Format Style",
		Items:     cmStyles,
		Templates: templates,
		Size:      6,
		Searcher:  searcher,
		CursorPos: styleIdx,
	}

	i, _, err := prompt.Run()
	utils.CheckError(err)

	if cmStyles[i].Name == "Quit" {
		os.Exit(1)
	}

	return cmStyles[i].CommitMessageStyle
}

func GetDefaultSelectTemplates() *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "➣ {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "➣ {{ .Name | red | cyan }}",
		Details:  config.DetailTmpl,
	}
}
