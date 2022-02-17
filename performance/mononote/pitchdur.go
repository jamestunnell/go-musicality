package mononote

import (
	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/performance/centpitch"
)

type PitchDur struct {
	Duration rat.Rat
	Pitch    *centpitch.CentPitch
}

func NewPitchDur(p *centpitch.CentPitch, dur rat.Rat) *PitchDur {
	return &PitchDur{
		Pitch:    p,
		Duration: dur,
	}
}
