package validation

func VerifyNonZeroUInt64(name string, val uint64) error {
	if val == 0 {
		return NewErrZero(name)
	}

	return nil
}
