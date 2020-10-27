package goexe

import (
	"os"
	"os/exec"
	"syscall"
)

const (
	EXIT_OK  = 0 // Zero exit code
	EXIT_BAD = 1 // General error exit code
)

// Run runs the executable at 'exePath'. Setting the 'workDir' as empty will
// use the default as specified by functions that accept exec.Cmd. EXIT_OK is
// returned on successful execution otherwise EXIT_BAD or another non-zero
// exit code is returned.
func RunCmd(exePath string, workDir string, args ...string) (int, error) {

	var e error

	cmd := exec.Command(exePath, args...)
	cmd.Dir = workDir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if e := cmd.Start(); e != nil {
		return EXIT_BAD, e
	}

	if e = cmd.Wait(); e == nil {
		return EXIT_OK, nil
	}

	if exitErr, ok := e.(*exec.ExitError); ok {
		if stat, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			return stat.ExitStatus(), e
		}
	}

	return EXIT_BAD, e
}
