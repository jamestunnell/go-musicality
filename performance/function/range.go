package function

import "math/big"

type Range struct {
	Start, End *big.Rat
}

func NewRange(start, end *big.Rat) Range {
	return Range{Start: start, End: end}
}

func (r Range) IsValid() bool {
	return r.End.Cmp(r.Start) == 1
}

func (r Range) Span() *big.Rat {
	return new(big.Rat).Sub(r.End, r.Start)
}

// IncludesRange checks to see if the current range includes the given range.
func (r Range) IncludesRange(r2 Range) bool {
	// Testing (r2.Start >= r.Start) && (r2.Start <= r.End) && (r2.End >= r.Start) && (r2.End <= r.End)
	return (r2.Start.Cmp(r.Start) >= 0) &&
		(r2.Start.Cmp(r.End) <= 0) &&
		(r2.End.Cmp(r.Start) >= 0) &&
		(r2.End.Cmp(r.End) <= 0)
}

// IncludesValue checks to see if the current range includes the given value (includes end).
func (r Range) IncludesValue(val *big.Rat) bool {
	return (val.Cmp(r.Start) >= 0) && (val.Cmp(r.End) <= 0)
}

// IncludesValueExcl checks to see if the current range includes the given value (excludes end).
func (r Range) IncludesValueExcl(val *big.Rat) bool {
	return (val.Cmp(r.Start) >= 0) && (val.Cmp(r.End) < 0)
}
