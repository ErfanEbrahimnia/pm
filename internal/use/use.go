package use

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ErfanEbrahimnia/pm/internal/manager"
)

// Run interactively sets package.json's packageManager field for the project containing startDir.
func Run(startDir string) error {
	m, err := selectManager()
	if err != nil {
		return err
	}

	projectDir, err := findProjectDir(startDir)
	if err != nil {
		return err
	}

	path := filepath.Join(projectDir, "package.json")
	if err := setPackageManager(path, m, ""); err != nil {
		return err
	}

	fmt.Printf("set packageManager to %q in %s\n", manager.PackageManagerField(m), path)
	return nil
}

func findProjectDir(startDir string) (string, error) {
	dir, err := filepath.Abs(startDir)
	if err != nil {
		return "", fmt.Errorf("resolve directory: %w", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "package.json")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("no package.json found from %s", startDir)
		}
		dir = parent
	}
}

func setPackageManager(path string, m manager.Manager, version string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read package.json: %w", err)
	}

	var pkg map[string]json.RawMessage
	if err := json.Unmarshal(data, &pkg); err != nil {
		return fmt.Errorf("parse package.json: %w", err)
	}

	value, err := json.Marshal(manager.PackageManagerFieldWithVersion(m, version))
	if err != nil {
		return err
	}
	pkg["packageManager"] = value

	out, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return fmt.Errorf("encode package.json: %w", err)
	}
	out = append(out, '\n')

	if err := os.WriteFile(path, out, 0o644); err != nil {
		return fmt.Errorf("write package.json: %w", err)
	}
	return nil
}
