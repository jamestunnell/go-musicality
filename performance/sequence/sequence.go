package sequence

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/duration"
)

// Sequence is a sequence of elements, starting at an offset, and with some
// kind of separation at the end.
type Sequence struct {
	Start    *big.Rat
	Elements []*Element
}

func New(start *big.Rat, elements ...*Element) *Sequence {
	return &Sequence{
		Start:    start,
		Elements: elements,
	}
}

func (s *Sequence) Validate() error {
	for _, elem := range s.Elements {
		if err := elem.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (seq *Sequence) Offsets() []*big.Rat {
	offsets := make([]*big.Rat, len(seq.Elements))
	currentOffset := new(big.Rat).Set(seq.Start)

	for i, e := range seq.Elements {
		offsets[i] = new(big.Rat).Set(currentOffset)
		currentOffset = new(big.Rat).Add(currentOffset, e.Duration.Rat)
	}

	return offsets
}

func (seq *Sequence) Duration() *big.Rat {
	dur := duration.Zero()

	for _, e := range seq.Elements {
		dur.Add(e.Duration)
	}

	return dur.Rat
}

func (seq *Sequence) End() *big.Rat {
	return new(big.Rat).Add(seq.Start, seq.Duration())
}

func (seq *Sequence) Simplify() error {
	if err := seq.Validate(); err != nil {
		return err
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
			prev.Duration.Add(cur.Duration)

			seq.Elements = append(seq.Elements[:i], seq.Elements[i+1:]...)
		} else {
			i++
		}
	}

	return nil
}
