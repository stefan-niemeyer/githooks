package styles

import (
	"testing"
)

func TestParseFormatStyle(t *testing.T) {
	// table w/ all test cases for ParseFormatStyle
	tests := []struct {
		name    string
		input   string
		want    FormatStyle
		wantErr bool
	}{
		{"Valid: PlainSpace", "plain_space", PlainSpace, false},
		{"Valid: PlainColon", "plain_colon", PlainColon, false},
		{"Valid: PlainBracket", "plain_brackets", PlainBrackets, false},
		{"Valid: Conventional", "conventional", Conventional, false},
		{"Empty String becomes PlainBrackets", "", PlainBrackets, false},
		{"Invalid Style", "something_different", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFormatStyle(tt.input)

			// check for wrong errors
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFormatStyle(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}

			// check the result
			if got != tt.want {
				t.Errorf("ParseFormatStyle(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestAddTicket(t *testing.T) {
	ticket := "JIRA-123"

	// table w/ all test cases for AddTicket
	tests := []struct {
		name  string
		msg   string
		style FormatStyle
		want  string
	}{
		// Plain variants
		{"PlainSpace: normal word", "add feature foo", PlainSpace, "JIRA-123 add feature foo"},
		{"PlainColon: normal word", "add feature foo", PlainColon, "JIRA-123: add feature foo"},
		{"PlainBracket: normal word", "add feature foo", PlainBrackets, "[JIRA-123] add feature foo"},

		// Special test case / fallback to default PlainBrackets
		{"Unknown style (fallback)", "add feature foo", FormatStyle("unknown"), "[JIRA-123] add feature foo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddTicket("", tt.msg, ticket, tt.style)
			if got != tt.want {
				t.Errorf("AddTicket(%q, %q, %v) = %v, want %v", tt.msg, ticket, tt.style, got, tt.want)
			}
		})
	}
}
