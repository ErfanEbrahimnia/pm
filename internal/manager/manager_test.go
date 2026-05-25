package manager_test

import (
	"testing"

	"github.com/ErfanEbrahimnia/pm/internal/manager"
)

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
