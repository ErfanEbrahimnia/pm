package main

import (
	"fmt"
	"os"

	"github.com/ErfanEbrahimnia/pm/internal/detect"
	"github.com/ErfanEbrahimnia/pm/internal/runner"
	"github.com/ErfanEbrahimnia/pm/internal/use"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "pm: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) > 0 && args[0] == "use" {
		if len(args) > 1 {
			return fmt.Errorf("pm use takes no arguments; run pm use to choose a package manager")
		}
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		return use.Run(cwd)
	}

	m, projectDir, err := detect.FromCwd()
	if err != nil {
		return err
	}

	if len(args) == 0 {
		fmt.Printf("package manager: %s (project: %s)\n", m, projectDir)
		return nil
	}

	return runner.Exec(m, args)
}
