package detect_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ErfanEbrahimnia/pm/internal/detect"
	"github.com/ErfanEbrahimnia/pm/internal/manager"
)

func TestFromDir_lockfilePriority(t *testing.T) {
	root := t.TempDir()
	sub := filepath.Join(root, "apps", "web")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatal(err)
	}

	write(t, root, "package-lock.json", "{}")
	write(t, root, "pnpm-lock.yaml", "lockfileVersion: 5.4\n")

	got, dir, err := detect.FromDir(sub)
	if err != nil {
		t.Fatal(err)
	}
	if got != manager.Pnpm {
		t.Fatalf("got %s, want pnpm", got)
	}
	if dir != root {
		t.Fatalf("project dir %q, want %q", dir, root)
	}
}

func TestFromDir_walksUp(t *testing.T) {
	root := t.TempDir()
	sub := filepath.Join(root, "nested")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	write(t, root, "yarn.lock", "# yarn lockfile\n")

	got, _, err := detect.FromDir(sub)
	if err != nil {
		t.Fatal(err)
	}
	if got != manager.Yarn {
		t.Fatalf("got %s, want yarn", got)
	}
}

func TestFromDir_packageManagerField(t *testing.T) {
	root := t.TempDir()
	write(t, root, "package.json", `{"name":"demo","packageManager":"bun@1.2.3"}`)
	write(t, root, "pnpm-lock.yaml", "lockfileVersion: 5.4\n")

	got, _, err := detect.FromDir(root)
	if err != nil {
		t.Fatal(err)
	}
	if got != manager.Bun {
		t.Fatalf("got %s, want bun from packageManager field", got)
	}
}

func TestFromDir_notFound(t *testing.T) {
	root := t.TempDir()
	_, _, err := detect.FromDir(root)
	if err == nil {
		t.Fatal("expected error when no lockfile or packageManager")
	}
	msg := err.Error()
	for _, name := range []string{"pnpm", "bun", "yarn", "npm"} {
		if !strings.Contains(msg, name) {
			t.Fatalf("error %q should mention supported manager %q", msg, name)
		}
	}
}

func write(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
