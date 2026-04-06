package hooks

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/stefan-niemeyer/githooks/config"
	"github.com/stefan-niemeyer/githooks/styles"
	"github.com/stefan-niemeyer/githooks/utils"
)

var conventionalRE = regexp.MustCompile(`^([a-zA-Z]+)(?:\([^)]*\))?!?:.*`)

func CommitMsg(filename string) {
	msgRaw, err := os.ReadFile(filename) // Read the commit message
	utils.CheckError(err)
	msg := string(msgRaw)

	// Don't touch merge  commits
	if strings.HasPrefix(msg, "Merge") {
		os.Exit(0)
	}

	// read Jira projects regex from Git config
	// read reference where HEAD is pointing to (e.g. branch name)
	jiraProjects := utils.GetGitConfigValue("user.jiraProjects")
	ticketRE := reFromProjects(jiraProjects)

	branch := utils.GetCurrentBranch()
	branchTicket := strings.ToUpper(ticketRE.FindString(branch))

	commitMessageStyle := utils.GetGitConfigValue("user.commitMessageStyle")
	cmStyle, errStyles := styles.ParseFormatStyle(commitMessageStyle)
	utils.CheckError(errStyles)

	if cmStyle == styles.Conventional {
		// check if commit message follows conventional commit format
		allowedTypes := utils.GetGitConfigValue("user.allowedTypes")
		if len(allowedTypes) == 0 {
			allowedTypes = config.AllowedTypesDefault
		}

		matches := conventionalRE.FindStringSubmatch(msg)
		if len(matches) < 2 {
			_, err = fmt.Fprintf(os.Stderr, "Commit message must follow convention and start with one of:\n  %s\n\nMessage:\n%s", allowedTypes, msg)
			os.Exit(1)
		}

		if !strings.Contains(","+allowedTypes+",", ","+matches[1]+",") {
			// format is OK, but the 'type' is not allowedTypes
			_, err = fmt.Fprintf(os.Stderr, "Commit message must start with one of:\n  %s\n", allowedTypes)
			os.Exit(1)
		}
	}

	// Check if the branch contains a valid issue that does not appear in the commit message.
	if len(branchTicket) != 0 {
		reBranchTicket := regexp.MustCompile("\\b" + branchTicket + "\\b")
		if !reBranchTicket.MatchString(msg) {
			// the ticket from the branch is not in the commit message yet
			msg = styles.AddTicket(filename, msg, branchTicket, cmStyle)
			fmt.Printf("New commit message:\n%s\n", msg)
		}
	}

	cmTicket := ticketRE.FindString(msg)
	// Check if commit message contains valid issue keys?
	if cmTicket == "" {
		if len(jiraProjects) != 0 {
			_, err = fmt.Fprintf(os.Stderr, "ERROR: The commit message is missing a Jira issue key with a project key matching '%s' (e.g. DS-123)\n", jiraProjects)
		} else {
			_, err = fmt.Fprintf(os.Stderr, "ERROR: The commit message is missing a Jira issue key (e.g. DS-123).\n")
		}
		os.Exit(1)
	}

	os.Exit(0)
}

func reFromProjects(jiraProjects string) *regexp.Regexp {
	if len(jiraProjects) == 0 {
		return regexp.MustCompile("\\b[[:alpha:]][[:alnum:]]*-[[:digit:]]+\\b")
	}
	return regexp.MustCompile("\\b" + jiraProjects + "-[[:digit:]]+\\b")
}
