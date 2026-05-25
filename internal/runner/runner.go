package runner

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ErfanEbrahimnia/pm/internal/manager"
)

// Exec runs the given package manager with args, inheriting stdio.
func Exec(m manager.Manager, args []string) error {
	if m == manager.Unknown {
		return fmt.Errorf("cannot run unknown package manager")
	}

	path, err := exec.LookPath(m.Command())
	if err != nil {
		return fmt.Errorf("%s is not installed or not on PATH: %w", m.Command(), err)
	}

	cmd := exec.Command(path, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		if exit, ok := err.(*exec.ExitError); ok {
			os.Exit(exit.ExitCode())
		}
		return err
	}
	return nil
}
