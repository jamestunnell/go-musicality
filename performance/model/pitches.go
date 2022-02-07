package model

import "github.com/jamestunnell/go-musicality/notation/rat"

type Pitches []*Pitch

func (ps Pitches) MakePitchDurs(dur rat.Rat) []*PitchDur {
	n := len(ps)
	pds := make([]*PitchDur, n)

	for i := 0; i < n; i++ {
		pds[i] = NewPitchDur(ps[i], dur)
	}

	return pds
}
