package model

import (
	"math/big"
)

type Note struct {
	Start      *big.Rat
	PitchDurs  []*PitchDur
	Attack     float64
	Separation float64
}

func NewNote(start *big.Rat, pitchDurs ...*PitchDur) *Note {
	return &Note{
		Start:      start,
		Attack:     0.0,
		Separation: 0.0,
		PitchDurs:  pitchDurs,
	}
}

func (seq *Note) Offsets() []*big.Rat {
	offsets := make([]*big.Rat, len(seq.PitchDurs))
	currentOffset := new(big.Rat).Set(seq.Start)

	for i, e := range seq.PitchDurs {
		offsets[i] = currentOffset
		currentOffset = new(big.Rat).Add(currentOffset, e.Duration)
	}

	return offsets
}

// Duration is not modified to account for Note separation.
func (seq *Note) Duration() *big.Rat {
	dur := big.NewRat(0, 1)

	for _, pd := range seq.PitchDurs {
		dur = dur.Add(dur, pd.Duration)
	}

	return dur
}

// End is not modified to account for separation
func (seq *Note) End() *big.Rat {
	end := new(big.Rat).Add(seq.Start, seq.Duration())

	return end
}

func (seq *Note) Simplify() {
	i := 1

	for i < len(seq.PitchDurs) {
		cur := seq.PitchDurs[i]
		prev := seq.PitchDurs[i-1]

		if cur.Pitch.Equal(prev.Pitch) {
			// combine current with previous element
			prev.Duration = prev.Duration.Add(prev.Duration, cur.Duration)

			seq.PitchDurs = append(seq.PitchDurs[:i], seq.PitchDurs[i+1:]...)
		} else {
			i++
		}
	}
}
