package validation

// Validatable can run validation checks and capture each failure individually.
type Validatable interface {
	// Run validation checks and capture each failure individually.
	// Return nil if there are no failures.
	Validate() *Result
}
