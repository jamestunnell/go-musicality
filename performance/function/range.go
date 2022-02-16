package function

import (
	"github.com/jamestunnell/go-musicality/notation/rat"
)

type Range struct {
	Start, End rat.Rat
}

func NewRange(start, end rat.Rat) Range {
	return Range{Start: start, End: end}
}

func (r Range) IsValid() bool {
	return r.End.Greater(r.Start)
}

func (r Range) Span() rat.Rat {
	return r.End.Sub(r.Start)
}

// IncludesRange checks to see if the current range includes the given range.
func (r Range) IncludesRange(r2 Range) bool {
	return r2.Start.GreaterEqual(r.Start) &&
		r2.Start.LessEqual(r.End) &&
		r2.End.GreaterEqual(r.Start) &&
		r2.End.LessEqual(r.End)
}

// IncludesValue checks to see if the current range includes the given value (includes end).
func (r Range) IncludesValue(val rat.Rat) bool {
	return val.GreaterEqual(r.Start) && val.LessEqual(r.End)
}

// IncludesValueExcl checks to see if the current range includes the given value (excludes end).
func (r Range) IncludesValueExcl(val rat.Rat) bool {
	return val.GreaterEqual(r.Start) && val.Less(r.End)
}
