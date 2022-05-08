package types

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	. "github.com/xiabai84/githooks/utils"
)

type Dialog struct {
	ErrorMsg string
	Label    string
}

func GetPromptInput(pc Dialog) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
		}
		return nil
	}
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . | }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}
	result, err := prompt.Run()
	CheckError(err)
	fmt.Printf("✅  %s\n", result)
	return result
}

func GetPromptSelect(pc Dialog, items []string) string {
	index := -1
	var result string
	var err error

	for index < 0 {

		prompt := promptui.SelectWithAdd{
			Label: pc.Label,
			Items: items,
		}
		index, result, err = prompt.Run()
		if index == -1 {
			items = append(items, result)
		}
	}
	CheckError(err)
	fmt.Printf("✅  Selected: %s\n", result)
	return result
}
