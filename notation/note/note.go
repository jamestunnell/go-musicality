package note

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/pkg/util"
)

const (
	Legato        = "legato"
	Tenuto        = "tenuto"
	Accent        = "accent"
	Marcato       = "marcato"
	Portato       = "portato"
	Staccato      = "staccato"
	Staccatissimo = "staccatissimo"
)

type Note struct {
	Pitches      []*pitch.Pitch `json:"pitches,omitempty"`
	Duration     *big.Rat       `json:"duration"`
	Articulation string         `json:"articulation,omitempty"`
}

func New(dur *big.Rat, pitches ...*pitch.Pitch) *Note {
	return &Note{
		Pitches:      pitches,
		Duration:     dur,
		Articulation: "",
	}
}

func (n *Note) Validate() error {
	if n.Duration.Cmp(big.NewRat(0, 1)) != 1 {
		return util.NewNonPositiveRatError(n.Duration)
	}

	switch n.Articulation {
	case "", Legato, Tenuto, Accent, Marcato, Portato, Staccato, Staccatissimo:
		// do nothing
	default:
		return fmt.Errorf("unknown articulation %s", n.Articulation)
	}

	return nil
}

func (n *Note) Dot() {
	n.Duration = new(big.Rat).Mul(n.Duration, big.NewRat(3, 2))
}

func (n *Note) IsRest() bool {
	return len(n.Pitches) == 0
}

func (n *Note) IsMonophonic() bool {
	return len(n.Pitches) == 1
}
