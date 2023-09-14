package function

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/common/rat"
)

type Range struct {
	Start, End *big.Rat
}

func NewRange(start, end *big.Rat) Range {
	return Range{Start: start, End: end}
}

func (r Range) IsValid() bool {
	return rat.IsGreater(r.End, r.Start)
}

func (r Range) Span() *big.Rat {
	return rat.Sub(r.End, r.Start)
}

// IncludesRange checks to see if the current range includes the given range.
func (r Range) IncludesRange(r2 Range) bool {
	return rat.IsGreaterEqual(r2.Start, r.Start) &&
		rat.IsLessEqual(r2.Start, r.End) &&
		rat.IsGreaterEqual(r2.End, r.Start) &&
		rat.IsLessEqual(r2.End, r.End)
}

// IncludesValue checks to see if the current range includes the given value (includes end).
func (r Range) IncludesValue(val *big.Rat) bool {
	return rat.IsGreaterEqual(val, r.Start) && rat.IsLessEqual(val, r.End)
}

// IncludesValueExcl checks to see if the current range includes the given value (excludes end).
func (r Range) IncludesValueExcl(val *big.Rat) bool {
	return rat.IsGreaterEqual(val, r.Start) && rat.IsLess(val, r.End)
}
