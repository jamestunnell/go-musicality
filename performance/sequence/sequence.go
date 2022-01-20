package sequence

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/validation"
)

// Sequence is a sequence of elements, starting at an offset, and with some
// kind of separation at the end.
type Sequence struct {
	Start      *big.Rat
	Elements   []*Element
	Separation float64
}

var errInvalidSeq = errors.New("sequence is invalid, validate for details")

func New(start *big.Rat, elements ...*Element) *Sequence {
	return &Sequence{
		Start:      start,
		Elements:   elements,
		Separation: SeparationNormal,
	}
}

func (s *Sequence) Validate() *validation.Result {
	results := []*validation.Result{}

	for i, elem := range s.Elements {
		if result := elem.Validate(); result != nil {
			result.Context = fmt.Sprintf("%s %d", result.Context, i)

			results = append(results, result)
		}
	}

	if len(results) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "sequence",
		Errors:     []error{},
		SubResults: results,
	}
}

func (seq *Sequence) Offsets() []*big.Rat {
	offsets := make([]*big.Rat, len(seq.Elements))
	currentOffset := new(big.Rat).Set(seq.Start)

	for i, e := range seq.Elements {
		offsets[i] = currentOffset
		currentOffset = new(big.Rat).Add(currentOffset, e.Duration)
	}

	return offsets
}

// Duration is not modified to account for sequence separation.
func (seq *Sequence) Duration() *big.Rat {
	dur := big.NewRat(0, 1)

	for _, e := range seq.Elements {
		dur = dur.Add(dur, e.Duration)
	}

	return dur
}

// End is not modified to account for separation
func (seq *Sequence) End() *big.Rat {
	end := new(big.Rat).Add(seq.Start, seq.Duration())

	return end
}

func (seq *Sequence) Simplify() error {
	if result := seq.Validate(); result != nil {
		return errInvalidSeq
	}

	if len(seq.Elements) == 0 {
		return nil
	}

	i := 1

	for i < len(seq.Elements) {
		cur := seq.Elements[i]
		prev := seq.Elements[i-1]

		if (cur.Pitch == prev.Pitch) && (cur.Attack == 0.0) {
			// combine current with previous element
			prev.Duration = prev.Duration.Add(prev.Duration, cur.Duration)

			seq.Elements = append(seq.Elements[:i], seq.Elements[i+1:]...)
		} else {
			i++
		}
	}

	return nil
}
