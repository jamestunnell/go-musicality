package performance

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/pkg/notation"
)

type PitchDur struct {
	Pitch    *notation.Pitch
	Duration *big.Rat
}

func NewPitchDur(p *notation.Pitch, dur *big.Rat) (*PitchDur, error) {
	if dur.Cmp(big.NewRat(0, 1)) <= 1 {
		return nil, fmt.Errorf("duration %v is non-positive", dur)
	}

	return &PitchDur{Pitch: p, Duration: dur}, nil
}
