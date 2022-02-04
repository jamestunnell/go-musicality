package function

type Range struct {
	Start float64
	End   float64
}

func NewRange(start, end float64) Range {
	return Range{Start: start, End: end}
}

func (r Range) IsValid() bool {
	return r.End > r.Start
}

func (r Range) Span() float64 {
	return r.End - r.Start
}

// IncludesRange checks to see if the current range includes the given range.
func (r Range) IncludesRange(r2 Range) bool {
	return (r2.Start >= r.Start) && (r2.Start <= r.End) && (r2.End >= r.Start) && (r2.End <= r.End)
}

// IncludesValue checks to see if the current range includes the given value (includes end).
func (r Range) IncludesValue(val float64) bool {
	return (val >= r.Start) && (val <= r.End)
}

// IncludesValueExcl checks to see if the current range includes the given value (excludes end).
func (r Range) IncludesValueExcl(val float64) bool {
	return (val >= r.Start) && (val < r.End)
}
