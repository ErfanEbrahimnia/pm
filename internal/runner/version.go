package runner

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/ErfanEbrahimnia/pm/internal/manager"
)

// Version runs `<manager> --version` and returns the trimmed output.
func Version(m manager.Manager) (string, error) {
	if m == manager.Unknown {
		return "", fmt.Errorf("cannot get version for unknown package manager")
	}

	path, err := exec.LookPath(m.Command())
	if err != nil {
		return "", fmt.Errorf("%s is not installed or not on PATH: %w", m.Command(), err)
	}

	cmd := exec.Command(path, "--version")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%s --version: %w", m.Command(), err)
	}

	out := strings.TrimSpace(stdout.String())
	if out == "" {
		out = strings.TrimSpace(stderr.String())
	}
	if out == "" {
		return "", fmt.Errorf("%s --version produced no output", m.Command())
	}

	// Use the first line (yarn can print extra lines).
	return strings.Split(out, "\n")[0], nil
}
