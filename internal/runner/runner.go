package runner

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

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

// ExecX runs the package manager's one-off execution tool (npx, pnpx, bunx, yarn dlx).
func ExecX(m manager.Manager, args []string) error {
	spec, err := m.ExecSpec()
	if err != nil {
		return err
	}

	path, err := exec.LookPath(spec.Command)
	if err != nil {
		return fmt.Errorf("%s is not installed or not on PATH: %w", spec.Command, err)
	}

	cmdArgs := append(append([]string{}, spec.Prefix...), args...)
	cmd := exec.Command(path, cmdArgs...)
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

// ExecVersion runs the detected exec tool's --version (npx, pnpx, bunx, yarn).
func ExecVersion(m manager.Manager) (string, error) {
	spec, err := m.ExecSpec()
	if err != nil {
		return "", err
	}

	path, err := exec.LookPath(spec.Command)
	if err != nil {
		return "", fmt.Errorf("%s is not installed or not on PATH: %w", spec.Command, err)
	}

	args := append(append([]string{}, spec.Prefix...), "--version")
	cmd := exec.Command(path, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%s --version: %w", formatExecLabel(spec), err)
	}

	out := strings.TrimSpace(stdout.String())
	if out == "" {
		out = strings.TrimSpace(stderr.String())
	}
	if out == "" {
		return "", fmt.Errorf("%s --version produced no output", formatExecLabel(spec))
	}

	return strings.Split(out, "\n")[0], nil
}

func formatExecLabel(spec manager.ExecSpec) string {
	if len(spec.Prefix) == 0 {
		return spec.Command
	}
	return spec.Command + " " + strings.Join(spec.Prefix, " ")
}
