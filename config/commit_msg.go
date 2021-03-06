package config

var CommitMsg = `#!/usr/bin/env bash

# List of Jira projects whose keys will be accepted by the hook.
# The installation procedure creates a RegExp of a comma separated list of project keys.
# PROJECTS="(DS|MYJIRA|MARS)"               # Accept issue keys like DS-17 or MARS-6
# PROJECTS="(MYJIRA)"                       # Accept only issue keys that start with MYJIRA, like MYJIRA-1966
# PROJECTS="MYJIRA"                         # Same as "(MYJIRA)"
# PROJECTS=""                               # Accept every issue key matching [[:alpha:]][[:alnum:]]*-[[:digit:]]+

PROJECTS=""

if [[ "$PROJECTS" == "" ]]; then
  # find if and where user.jiraProjects is set with
  # git config --show-origin user.jiraProjects
  PROJECTS=$(git config --get user.jiraProjects)
fi

# Add git branch if relevant
parse_git_branch() {
  if [ -n "$PROJECTS" ]; then
    # Parse the current branch name for an issue where the project key matches the regular expression in PROJECTS
    git rev-parse --abbrev-ref HEAD 2>/dev/null | \
        grep --ignore-case --extended-regexp --regexp="\<${PROJECTS}-[[:digit:]]+\>" --only-matching | \
        tr '[:lower:]' '[:upper:]'
  else
    # Parse the current branch for any issue key
    git rev-parse --abbrev-ref HEAD 2>/dev/null | \
        tr '[:lower:]' '[:upper:]'
  fi
}

# Extact ticket number (e.g. DS-123) from first line of the commit message
parse_first_message_line_for_tickets() {    # Pass the name of a file with the commit message as 1st parameter
  if [ -n "$PROJECTS" ]; then
    # Parse the first line of the commit message for an issue where the project key matches the regular expression in PROJECTS
    echo "$FIRST_LINE" | \
        grep --ignore-case --extended-regexp --regexp="\<${PROJECTS}-[[:digit:]]+\>" --only-matching | \
        tr '[:lower:]' '[:upper:]'
  else
    # Parse the first line of the commit message for any issue key
    echo "$FIRST_LINE" | \
        grep --extended-regexp --regexp='\<[[:alpha:]][[:alnum:]]*-[[:digit:]]+\>' --only-matching | \
        tr '[:lower:]' '[:upper:]'
  fi
}

FIRST_LINE=$(head --lines=1 "$1")
BRANCH_TICKET=$(parse_git_branch)
CM_TICKETS=$(parse_first_message_line_for_tickets "$1")

# Check if the branch contains a valid issue that does not appear in the 1st line of the commit message.
if [[ "$BRANCH_TICKET" != "" && ! "$CM_TICKETS" =~ $BRANCH_TICKET ]]; then
  MESSAGE=$(cat "$1")
  echo "New commit message: [$BRANCH_TICKET] $MESSAGE"
  echo "[$BRANCH_TICKET] $MESSAGE" >"$1"      # Let the commit message start with the issue key found in the branch
  exit 0
fi

# Check if commit message contains valid issue keys or 'Merge'?
if [[ "$CM_TICKETS" == "" && ! $FIRST_LINE =~ "Merge" ]]; then
  if [ -n "$PROJECTS" ]; then
    echo >&2 "ERROR: The 1st line of the commit message is missing a Jira issue key with a project key matching '$PROJECTS' (e.g. DS-123)"
  else
    echo >&2 "ERROR: The 1st line of the commit message is missing a Jira issue key (e.g. DS-123)."
  fi
  exit 1
fi
`
