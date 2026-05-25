package version

import (
	"fmt"
	"io"
	"os"
	"strings"

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

// PrintPmx writes pmx's version and, when inside a project, the detected exec tool version.
func PrintPmx(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "pmx %s\n", Version); err != nil {
		return err
	}

	m, _, err := detect.FromCwd()
	if err != nil {
		return nil
	}

	execVersion, err := runner.ExecVersion(m)
	if err != nil {
		return err
	}

	spec, err := m.ExecSpec()
	if err != nil {
		return err
	}

	label := spec.Command
	if len(spec.Prefix) > 0 {
		label += " " + strings.Join(spec.Prefix, " ")
	}

	_, err = fmt.Fprintf(w, "%s %s\n", label, execVersion)
	return err
}

// PrintPmxStdout is a convenience wrapper around PrintPmx(os.Stdout).
func PrintPmxStdout() error {
	return PrintPmx(os.Stdout)
}
