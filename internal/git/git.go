package git

import (
	"os/exec"
)

func runGitCommand(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func Status() (string, error) {
	return runGitCommand("status")
}

func Commit(message string) (string, error) {
	return runGitCommand("commit", "-m", message)
}

func Push() (string, error) {
	return runGitCommand("push")
}
