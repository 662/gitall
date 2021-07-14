package services

import "os/exec"

func ExecGitCommand(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	msg, err := cmd.CombinedOutput()
	cmd.Run()
	return string(msg), err
}
