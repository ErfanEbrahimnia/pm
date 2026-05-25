package detect

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ErfanEbrahimnia/pm/internal/manager"
)

// lockfile signals a package manager via its lockfile name.
type lockfile struct {
	name    string
	manager manager.Manager
}

// Lockfile priority when multiple lockfiles exist in the same directory.
var lockfiles = []lockfile{
	{name: "pnpm-lock.yaml", manager: manager.Pnpm},
	{name: "bun.lockb", manager: manager.Bun},
	{name: "bun.lock", manager: manager.Bun},
	{name: "yarn.lock", manager: manager.Yarn},
	{name: "package-lock.json", manager: manager.Npm},
}

// FromDir walks from startDir up to the filesystem root and returns the
// detected package manager for the nearest project directory.
func FromDir(startDir string) (manager.Manager, string, error) {
	dir, err := filepath.Abs(startDir)
	if err != nil {
		return manager.Unknown, "", fmt.Errorf("resolve directory: %w", err)
	}

	for {
		if m, ok := detectInDir(dir); ok {
			return m, dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return manager.Unknown, "", fmt.Errorf(
				"no package manager found from %s\nsupported: %s; run pm use to set one",
				startDir,
				manager.SupportedList(),
			)
		}
		dir = parent
	}
}

// FromCwd detects using the current working directory.
func FromCwd() (manager.Manager, string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return manager.Unknown, "", fmt.Errorf("get working directory: %w", err)
	}
	return FromDir(cwd)
}

func detectInDir(dir string) (manager.Manager, bool) {
	if m, ok := fromPackageJSON(dir); ok {
		return m, true
	}
	return fromLockfiles(dir)
}

func fromPackageJSON(dir string) (manager.Manager, bool) {
	path := filepath.Join(dir, "package.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return manager.Unknown, false
	}

	var pkg struct {
		PackageManager string `json:"packageManager"`
	}
	if err := json.Unmarshal(data, &pkg); err != nil || pkg.PackageManager == "" {
		return manager.Unknown, false
	}

	m, err := manager.ParsePackageManagerField(pkg.PackageManager)
	if err != nil {
		return manager.Unknown, false
	}
	return m, true
}

func fromLockfiles(dir string) (manager.Manager, bool) {
	for _, lf := range lockfiles {
		if _, err := os.Stat(filepath.Join(dir, lf.name)); err == nil {
			return lf.manager, true
		}
	}
	return manager.Unknown, false
}
