package performance

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/pitch"
)

type Element struct {
	Duration *big.Rat
	Pitch    *pitch.Pitch
	Attack   Attack
}

// Sequence is a sequence of elements, starting at an offset, and with some
// kind of separation at the end.
type Sequence struct {
	Start      *big.Rat
	Elements   []*Element
	Separation Separation
}

var (
	zero = big.NewRat(0, 1)
)

// NewSequence creates a new, simplified note sequence (combine elements that have the
// same pitch and have no attack between them).
func NewSequence(
	start *big.Rat, elems []*Element, sep Separation) (*Sequence, error) {
	if len(elems) == 0 {
		return nil, errors.New("no elements given")
	}

	for i, e := range elems {
		if e.Duration.Cmp(zero) == 0 {
			return nil, fmt.Errorf("element[%d] has zero duration", i)
		}
	}

	seq := &Sequence{Start: start, Elements: simplifyElements(elems), Separation: sep}

	return seq, nil
}

func (seq *Sequence) Offsets() []*big.Rat {
	offsets := make([]*big.Rat, len(seq.Elements))
	currentOffset := new(big.Rat).Set(seq.Start)

	for i, e := range seq.Elements {
		offsets[i] = new(big.Rat).Set(currentOffset)
		currentOffset.Add(currentOffset, e.Duration)
	}

	return offsets
}

func (seq *Sequence) Duration() *big.Rat {
	dur := big.NewRat(0, 1)

	for _, e := range seq.Elements {
		dur.Add(dur, e.Duration)
	}

	return dur
}

func (seq *Sequence) End() *big.Rat {
	return new(big.Rat).Add(seq.Start, seq.Duration())
}

func simplifyElements(elems []*Element) []*Element {
	var prev *Element
	newElems := make([]*Element, 0, 1)

	addNewElem := func(e *Element) {
		newElem := &Element{Duration: e.Duration, Pitch: e.Pitch, Attack: e.Attack}
		newElems = append(newElems, newElem)
		prev = newElem
	}

	addNewElem(elems[0])

	for i := 1; i < len(elems); i++ {
		e := elems[i]

		if (e.Pitch == prev.Pitch) && (e.Attack == Attack0) {
			prev.Duration.Add(prev.Duration, e.Duration)
		} else {
			addNewElem(e)
		}
	}

	return newElems
}
