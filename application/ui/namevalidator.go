package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func MakeNameValidator(m *ItemManager) fyne.StringValidator {
	return func(s string) error {
		if len(s) == 0 {
			return fmt.Errorf("name is empty")
		}

		if m.HasItem(s) {
			return fmt.Errorf("part '%s' already exists", s)
		}

		return nil
	}
}
