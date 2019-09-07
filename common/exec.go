package common

import "os/exec"

// Exec run special cmd, block until cmd exits
func Exec(cmd string, args ...string) ([]byte, error) {
	return exec.Command(cmd, args...).CombinedOutput()
}
