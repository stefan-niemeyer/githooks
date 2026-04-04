package styles

import (
	"fmt"
	"os"

	"github.com/stefan-niemeyer/githooks/utils"
)

type FormatStyle string

// The allowed values as they should be in the Config file
const (
	PlainSpace    FormatStyle = "plain_space"
	PlainColon    FormatStyle = "plain_colon"
	PlainBrackets FormatStyle = "plain_brackets"
	Conventional  FormatStyle = "conventional"
)

// ParseFormatStyle validates a string and converts it to a FormatStyle
func ParseFormatStyle(configValue string) (FormatStyle, error) {
	style := FormatStyle(configValue)
	if configValue == "" {
		style = PlainBrackets
	}
	switch style {
	case PlainSpace, PlainColon, PlainBrackets, Conventional:
		return style, nil
	default:
		return "", fmt.Errorf("unknown FormatStyle: '%s'", configValue)
	}
}

func AddTicket(filename, msg, ticket string, style FormatStyle) string {
	newMsg := ""

	switch style {
	case PlainSpace:
		newMsg = fmt.Sprintf("%s %s", ticket, msg)

	case PlainColon:
		newMsg = fmt.Sprintf("%s: %s", ticket, msg)

	case Conventional:
		if len(ticket) != 0 {
			trailerKey := "Refs"
			newMsg = utils.CallInterpretTrailers(trailerKey, ticket, filename)
		}

	default: // PlainBrackets is the default
		newMsg = fmt.Sprintf("[%s] %s", ticket, msg)
	}

	if len(filename) != 0 {
		err := os.WriteFile(filename, []byte(newMsg), 0644)
		utils.CheckError(err)
	}

	return newMsg
}
