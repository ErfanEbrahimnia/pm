package use

import (
	"errors"
	"fmt"
	"os"

	"github.com/ErfanEbrahimnia/pm/internal/manager"
	"github.com/manifoldco/promptui"
	"golang.org/x/term"
)

func selectManager() (manager.Manager, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return manager.Unknown, fmt.Errorf("stdin is not a terminal; run pm use from an interactive shell")
	}

	items := make([]string, len(manager.Supported))
	for i, m := range manager.Supported {
		items[i] = m.String()
	}

	prompt := promptui.Select{
		Label:        "Package manager",
		Items:        items,
		Size:         len(manager.Supported),
		HideSelected: true,
	}

	idx, _, err := prompt.Run()
	if err != nil {
		if errors.Is(err, promptui.ErrInterrupt) {
			return manager.Unknown, fmt.Errorf("selection cancelled")
		}
		return manager.Unknown, err
	}
	return manager.Supported[idx], nil
}
