package runner_test

import (
	"testing"

	"github.com/ErfanEbrahimnia/pm/internal/manager"
	"github.com/ErfanEbrahimnia/pm/internal/runner"
)

func TestVersion_unknownManager(t *testing.T) {
	_, err := runner.Version(manager.Unknown)
	if err == nil {
		t.Fatal("expected error for unknown manager")
	}
}
