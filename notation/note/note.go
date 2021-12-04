package note

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/pkg/util"
)

type Note struct {
	Pitches  []*pitch.Pitch `json:"pitches,omitempty"`
	Duration *big.Rat       `json:"duration"`
}

var (
	zero = big.NewRat(0, 1)
)

func New(dur *big.Rat, pitches ...*pitch.Pitch) *Note {
	return &Note{Pitches: pitches, Duration: dur}
}

func (n *Note) Validate() error {
	if n.Duration.Cmp(zero) <= 0 {
		return util.NewNonPositiveRatError(n.Duration)
	}

	return nil
}

func (n *Note) IsRest() bool {
	return len(n.Pitches) == 0
}

func (n *Note) IsMonophonic() bool {
	return len(n.Pitches) == 1
}
