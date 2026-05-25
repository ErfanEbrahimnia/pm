package version

import (
	"fmt"
	"io"
	"os"

	"github.com/ErfanEbrahimnia/pm/internal/detect"
	"github.com/ErfanEbrahimnia/pm/internal/runner"
)

// Print writes pm's version and, when inside a project, the detected package manager's version.
func Print(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "pm %s\n", Version); err != nil {
		return err
	}

	m, _, err := detect.FromCwd()
	if err != nil {
		return nil
	}

	pmVersion, err := runner.Version(m)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "%s %s\n", m, pmVersion)
	return err
}

// PrintStdout is a convenience wrapper around Print(os.Stdout).
func PrintStdout() error {
	return Print(os.Stdout)
}
