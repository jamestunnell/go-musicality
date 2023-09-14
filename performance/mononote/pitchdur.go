package mononote

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/performance/centpitch"
)

type PitchDur struct {
	Duration *big.Rat
	Pitch    *centpitch.CentPitch
}

func NewPitchDur(p *centpitch.CentPitch, dur *big.Rat) *PitchDur {
	return &PitchDur{
		Pitch:    p,
		Duration: dur,
	}
}
