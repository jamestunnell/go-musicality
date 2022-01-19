package note

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/validation"
)

const (
	// Articulations
	Legato        = "legato"
	Tenuto        = "tenuto"
	Accent        = "accent"
	Marcato       = "marcato"
	Portato       = "portato"
	Staccato      = "staccato"
	Staccatissimo = "staccatissimo"
)

type Note struct {
	Pitches      *pitch.Set             `json:"pitches,omitempty"`
	Duration     *big.Rat               `json:"duration"`
	Articulation string                 `json:"articulation,omitempty"`
	Slurs        bool                   `json:"slurs,omitempty"`
	Links        map[*pitch.Pitch]*Link `json: "links,omitempty"`
}

func New(dur *big.Rat, pitches ...*pitch.Pitch) *Note {
	return &Note{
		Pitches:      pitch.NewSet(pitches...),
		Duration:     dur,
		Articulation: "",
		Slurs:        false,
		Links:        make(map[*pitch.Pitch]*Link),
	}
}

func (n *Note) Validate() *validation.Result {
	errs := []error{}

	if n.Duration.Cmp(big.NewRat(0, 1)) != 1 {
		errs = append(errs, validation.NewErrNonPositiveRat("duration", n.Duration))
	}

	switch n.Articulation {
	case "", Legato, Tenuto, Accent, Marcato, Portato, Staccato, Staccatissimo:
		// do nothing
	default:
		errs = append(errs, fmt.Errorf("unknown articulation %s", n.Articulation))
	}

	if len(errs) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "note",
		Errors:     errs,
		SubResults: []*validation.Result{},
	}
}

func (n *Note) Dot() {
	n.Duration = new(big.Rat).Mul(n.Duration, big.NewRat(3, 2))
}

func (n *Note) IsRest() bool {
	return n.Pitches.Len() == 0
}

func (n *Note) IsMonophonic() bool {
	return n.Pitches.Len() == 1
}
