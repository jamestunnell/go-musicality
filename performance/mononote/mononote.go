package mononote

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/common/rat"
)

// MonoNote is a note with only one pitch at a time.
// Pitch changes are meant to be immediate, with no separation.
type MonoNote struct {
	Start      *big.Rat
	PitchDurs  []*PitchDur
	Attack     float64
	Separation float64
}

// New makes a new MonoNote.
func New(start *big.Rat, pitchDurs ...*PitchDur) *MonoNote {
	return &MonoNote{
		Start:      start,
		Attack:     0.0,
		Separation: 0.0,
		PitchDurs:  pitchDurs,
	}
}

// Duration is the sum of the durations from PitchDurs.
// Not modified to account for Note separation.
func (n *MonoNote) Duration() *big.Rat {
	dur := rat.Zero()

	for _, pd := range n.PitchDurs {
		dur = rat.Add(dur, pd.Duration)
	}

	return dur
}

// End is not modified to account for separation
func (n *MonoNote) End() *big.Rat {
	return rat.Add(n.Start, n.Duration())
}

func (n *MonoNote) Simplify() {
	i := 1

	for i < len(n.PitchDurs) {
		cur := n.PitchDurs[i]
		prev := n.PitchDurs[i-1]

		if cur.Pitch.Equal(prev.Pitch) {
			// combine current with previous element
			prev.Duration = rat.Add(prev.Duration, cur.Duration)

			n.PitchDurs = append(n.PitchDurs[:i], n.PitchDurs[i+1:]...)
		} else {
			i++
		}
	}
}
