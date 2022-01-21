package validation

func VerifyInRangeFloat(name string, val, min, max float64) error {
	if val < min || val > max {
		return NewErrOutOfRange(name, val, min, max)
	}

	return nil
}
