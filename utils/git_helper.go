package utils

import (
	"os/exec"
	"strings"
)

func GetGitConfigValue(key string) string {
	// equivalent to: git config --get <key>
	cmd := exec.Command("git", "config", "--get", key)

	out, err := cmd.Output()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(out))
}

func GetCurrentBranch() string {
	// git rev-parse --abbrev-ref HEAD
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")

	out, err := cmd.Output()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(out))
}

func CallInterpretTrailers(trailerKey string, trailerValue string, filename string) string {
	// git interpret-trailers --in-place --trailer <trailerKey>=<trailerValue> --if-exists addIfDifferent
	cmd := exec.Command("git", "interpret-trailers",
		"--trailer="+trailerKey+"="+trailerValue,
		"--if-exists",
		"addIfDifferent",
		filename)

	out, err := cmd.Output()
	CheckError(err)

	return string(out)
}
