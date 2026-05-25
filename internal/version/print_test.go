package version_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ErfanEbrahimnia/pm/internal/version"
)

func TestPrint_devOnlyOutsideProject(t *testing.T) {
	version.Version = "1.0.0-test"

	dir := t.TempDir()
	chdir(t, dir)

	var out bytes.Buffer
	if err := version.Print(&out); err != nil {
		t.Fatal(err)
	}

	if got := out.String(); got != "pm 1.0.0-test\n" {
		t.Fatalf("got %q, want only pm version outside a project", got)
	}
}

func TestPrint_includesManagerWhenDetected(t *testing.T) {
	version.Version = "1.0.0-test"

	dir := t.TempDir()
	pkg := filepath.Join(dir, "package.json")
	if err := os.WriteFile(pkg, []byte(`{"packageManager":"npm@10.0.0"}`), 0o644); err != nil {
		t.Fatal(err)
	}
	chdir(t, dir)

	var out bytes.Buffer
	if err := version.Print(&out); err != nil {
		t.Skipf("npm not available for version check: %v", err)
	}

	got := out.String()
	if !strings.HasPrefix(got, "pm 1.0.0-test\n") {
		t.Fatalf("got %q, want pm version first", got)
	}
	if !strings.Contains(got, "npm ") {
		t.Fatalf("got %q, want npm version line", got)
	}
}

func chdir(t *testing.T, dir string) {
	t.Helper()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(cwd) })
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
}
