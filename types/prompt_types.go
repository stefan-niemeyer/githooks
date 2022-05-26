package types

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	. "github.com/stefan-niemeyer/githooks/utils"
)

type Dialog struct {
	ErrorMsg string
	Label    string
}

func GetPromptInput(pc Dialog, defaultInput string) string {
	validate := func(input string) error {
		if len(input) == 0 && len(defaultInput) == 0 {
			return errors.New(pc.ErrorMsg)
		}
		return nil
	}
	templates := &promptui.PromptTemplates{
		Prompt:  promptui.IconInitial + " {{ . | }} ",
		Valid:   promptui.IconGood + " {{ . | green }} ",
		Invalid: promptui.IconBad + " {{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}
	result, err := prompt.Run()
	CheckError(err)
	if len(result) == 0 {
		result = defaultInput
	}
	fmt.Printf(promptui.IconGood+"  %s\n", result)
	return result
}
