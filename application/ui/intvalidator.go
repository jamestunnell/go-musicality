package ui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
)

func MakeIntValidator(min, max int) fyne.StringValidator {
	return func(s string) error {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}

		if i < min || i > max {
			return fmt.Errorf("%d is not in range %d-%d", i, min, max)
		}

		return nil
	}
}
