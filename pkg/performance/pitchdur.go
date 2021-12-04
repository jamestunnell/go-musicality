package performance

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/pitch"
)

type PitchDur struct {
	Pitch    *pitch.Pitch
	Duration *big.Rat
}

func NewPitchDur(p *pitch.Pitch, dur *big.Rat) *PitchDur {
	return &PitchDur{Pitch: p, Duration: dur}
}

func (pd *PitchDur) Validate() error {
	if pd.Duration.Cmp(big.NewRat(0, 1)) <= 1 {
		return fmt.Errorf("duration %v is non-positive", pd.Duration)
	}

	return nil
}
