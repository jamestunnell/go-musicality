package notation

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/pkg/util"
)

type Note struct {
	Pitches  []*Pitch `json:"pitches,omitempty"`
	Duration *big.Rat `json:"duration"`
}

var (
	zero = big.NewRat(0, 1)
)

func NewNote(dur *big.Rat, pitches ...*Pitch) (*Note, error) {
	if dur.Cmp(zero) <= 0 {
		return nil, util.NewNonPositiveRatError(dur)
	}
	return &Note{Pitches: pitches, Duration: dur}, nil
}

func (n *Note) IsRest() bool {
	return len(n.Pitches) == 0
}

func (n *Note) IsMonophonic() bool {
	return len(n.Pitches) == 1
}
