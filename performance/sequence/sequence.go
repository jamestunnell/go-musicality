package sequence

import (
	"errors"
	"fmt"

	"github.com/jamestunnell/go-musicality/validation"
)

// Sequence is a sequence of elements, starting at an offset, and with some
// kind of separation at the end.
type Sequence struct {
	Start    float64
	Elements []*Element
}

var errInvalidSeq = errors.New("sequence is invalid, validate for details")

func New(start float64, elements ...*Element) *Sequence {
	return &Sequence{
		Start:    start,
		Elements: elements,
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

func (seq *Sequence) Offsets() []float64 {
	offsets := make([]float64, len(seq.Elements))
	currentOffset := seq.Start

	for i, e := range seq.Elements {
		offsets[i] = currentOffset
		currentOffset += e.Duration
	}

	return offsets
}

func (seq *Sequence) Duration() float64 {
	dur := 0.0

	for _, e := range seq.Elements {
		dur += e.Duration
	}

	return dur
}

func (seq *Sequence) End() float64 {
	return seq.Start + seq.Duration()
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
			prev.Duration += cur.Duration

			seq.Elements = append(seq.Elements[:i], seq.Elements[i+1:]...)
		} else {
			i++
		}
	}

	return nil
}
