package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ErfanEbrahimnia/pm/internal/detect"
	"github.com/ErfanEbrahimnia/pm/internal/manager"
	"github.com/ErfanEbrahimnia/pm/internal/runner"
	"github.com/ErfanEbrahimnia/pm/internal/version"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "pmx: %v\n", err)
		os.Exit(1)
	}
}

func isVersionFlag(arg string) bool {
	switch arg {
	case "--version", "-version", "-v":
		return true
	default:
		return false
	}
}

func run(args []string) error {
	if len(args) > 0 && isVersionFlag(args[0]) {
		return version.PrintPmxStdout()
	}

	m, projectDir, err := detect.FromCwd()
	if err != nil {
		return err
	}

	if len(args) == 0 {
		spec, err := m.ExecSpec()
		if err != nil {
			return err
		}
		fmt.Printf("exec: %s (package manager: %s, project: %s)\n", formatExec(spec), m, projectDir)
		return nil
	}

	return runner.ExecX(m, args)
}

func formatExec(spec manager.ExecSpec) string {
	if len(spec.Prefix) == 0 {
		return spec.Command
	}
	return spec.Command + " " + strings.Join(spec.Prefix, " ")
}
