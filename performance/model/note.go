package model

import (
	"github.com/jamestunnell/go-musicality/notation/rat"
)

type Note struct {
	Start      rat.Rat
	PitchDurs  []*PitchDur
	Attack     float64
	Separation float64
}

func NewNote(start rat.Rat, pitchDurs ...*PitchDur) *Note {
	return &Note{
		Start:      start,
		Attack:     0.0,
		Separation: 0.0,
		PitchDurs:  pitchDurs,
	}
}

func (n *Note) Offsets() rat.Rats {
	offsets := make(rat.Rats, len(n.PitchDurs))
	currentOffset := n.Start.Clone()

	for i, e := range n.PitchDurs {
		offsets[i] = currentOffset.Clone()
		currentOffset.Accum(e.Duration)
	}

	return offsets
}

// Duration is not modified to account for Note separation.
func (n *Note) Duration() rat.Rat {
	dur := rat.Zero()

	for _, pd := range n.PitchDurs {
		dur.Accum(pd.Duration)
	}

	return dur
}

// End is not modified to account for separation
func (n *Note) End() rat.Rat {
	return n.Start.Add(n.Duration())
}

func (n *Note) Simplify() {
	i := 1

	for i < len(n.PitchDurs) {
		cur := n.PitchDurs[i]
		prev := n.PitchDurs[i-1]

		if cur.Pitch.Equal(prev.Pitch) {
			// combine current with previous element
			prev.Duration.Accum(cur.Duration)

			n.PitchDurs = append(n.PitchDurs[:i], n.PitchDurs[i+1:]...)
		} else {
			i++
		}
	}
}
