package manager_test

import (
	"testing"

	"github.com/ErfanEbrahimnia/pm/internal/manager"
)

func TestExecSpec(t *testing.T) {
	tests := []struct {
		m    manager.Manager
		cmd  string
		prefix []string
	}{
		{manager.Npm, "npx", nil},
		{manager.Pnpm, "pnpx", nil},
		{manager.Bun, "bunx", nil},
		{manager.Yarn, "yarn", []string{"dlx"}},
	}

	for _, tc := range tests {
		spec, err := tc.m.ExecSpec()
		if err != nil {
			t.Fatalf("%s: %v", tc.m, err)
		}
		if spec.Command != tc.cmd {
			t.Fatalf("%s: command %q, want %q", tc.m, spec.Command, tc.cmd)
		}
		if len(spec.Prefix) != len(tc.prefix) {
			t.Fatalf("%s: prefix %v, want %v", tc.m, spec.Prefix, tc.prefix)
		}
		for i := range tc.prefix {
			if spec.Prefix[i] != tc.prefix[i] {
				t.Fatalf("%s: prefix[%d] %q, want %q", tc.m, i, spec.Prefix[i], tc.prefix[i])
			}
		}
	}
}

func TestParsePackageManagerField(t *testing.T) {
	tests := []struct {
		in   string
		want manager.Manager
	}{
		{"pnpm@9.0.0", manager.Pnpm},
		{"yarn@4.0.0", manager.Yarn},
		{"bun@1.0.0", manager.Bun},
		{"npm@10.0.0", manager.Npm},
	}

	for _, tc := range tests {
		got, err := manager.ParsePackageManagerField(tc.in)
		if err != nil {
			t.Fatalf("%q: %v", tc.in, err)
		}
		if got != tc.want {
			t.Fatalf("%q: got %s, want %s", tc.in, got, tc.want)
		}
	}
}
