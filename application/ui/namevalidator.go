package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func MakeNameValidator(m *ItemManager, existing Item) fyne.StringValidator {
	return func(s string) error {
		if len(s) == 0 {
			return fmt.Errorf("name is empty")
		}

		if existing != nil {
			if s != existing.Name() && m.HasItem(s) {
				return fmt.Errorf("name '%s' already exists", s)
			}
		} else if m.HasItem(s) {
			return fmt.Errorf("name '%s' already exists", s)
		}

		return nil
	}
}
