package ui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"golang.org/x/exp/constraints"
)

func MakeNumericValidator[T constraints.Ordered](
	conv func(string) (T, error), min, max T) fyne.StringValidator {
	return func(s string) error {
		numVal, err := conv(s)
		if err != nil {
			return err
		}

		if numVal < min || numVal > max {
			return fmt.Errorf("%v is not in range %v-%v", numVal, min, max)
		}

		return nil
	}
}

func MakeIntValidator(min, max int) fyne.StringValidator {
	return MakeNumericValidator(strconv.Atoi, min, max)
}

func MakeFloatValidator(min, max float64) fyne.StringValidator {
	return MakeNumericValidator(Atof, min, max)
}

func Atof(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
