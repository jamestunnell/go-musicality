package validation

func VerifyNonZeroUInt(name string, val uint) error {
	if val == 0 {
		return NewErrZero(name)
	}

	return nil
}
