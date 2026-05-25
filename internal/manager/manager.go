package manager

import (
	"fmt"
	"strings"
)

// Manager identifies a Node.js package manager.
type Manager int

const (
	Unknown Manager = iota
	Pnpm
	Bun
	Yarn
	Npm
)

var names = [...]string{
	Unknown: "unknown",
	Pnpm:    "pnpm",
	Bun:     "bun",
	Yarn:    "yarn",
	Npm:     "npm",
}

// Supported lists package managers pm can detect and run.
var Supported = []Manager{Pnpm, Bun, Yarn, Npm}

// SupportedList returns a comma-separated list of supported package manager names.
func SupportedList() string {
	names := make([]string, len(Supported))
	for i, m := range Supported {
		names[i] = m.String()
	}
	return strings.Join(names, ", ")
}

func (m Manager) String() string {
	if int(m) < 0 || int(m) >= len(names) {
		return names[Unknown]
	}
	return names[m]
}

// Command returns the executable name on PATH.
func (m Manager) Command() string {
	return m.String()
}

// ExecSpec describes how to run one-off package binaries (npx, pnpx, bunx, yarn dlx).
type ExecSpec struct {
	Command string
	Prefix  []string // prepended before user args (e.g. "dlx" for yarn)
}

// ExecSpec returns the command used to execute packages without a global install.
func (m Manager) ExecSpec() (ExecSpec, error) {
	switch m {
	case Npm:
		return ExecSpec{Command: "npx"}, nil
	case Pnpm:
		return ExecSpec{Command: "pnpx"}, nil
	case Bun:
		return ExecSpec{Command: "bunx"}, nil
	case Yarn:
		return ExecSpec{Command: "yarn", Prefix: []string{"dlx"}}, nil
	default:
		return ExecSpec{}, fmt.Errorf("unknown package manager")
	}
}

// Parse recognizes a manager name (e.g. "pnpm", "PNPM").
func Parse(s string) (Manager, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "pnpm":
		return Pnpm, nil
	case "bun":
		return Bun, nil
	case "yarn":
		return Yarn, nil
	case "npm":
		return Npm, nil
	default:
		return Unknown, fmt.Errorf("unknown package manager %q (want pnpm, bun, yarn, or npm)", s)
	}
}

// ParsePackageManagerField parses the package.json "packageManager" value (e.g. "pnpm@9.0.0").
func ParsePackageManagerField(value string) (Manager, error) {
	value = strings.TrimSpace(value)
	name, _, _ := strings.Cut(value, "@")
	if name == "" {
		return Unknown, fmt.Errorf("invalid packageManager field %q", value)
	}
	return Parse(name)
}

// PackageManagerField returns a package.json packageManager value without a version pin.
func PackageManagerField(m Manager) string {
	return m.String() + "@"
}

// PackageManagerFieldWithVersion returns a fully qualified packageManager value.
func PackageManagerFieldWithVersion(m Manager, version string) string {
	if version == "" {
		return PackageManagerField(m)
	}
	return m.String() + "@" + version
}
