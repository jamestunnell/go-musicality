package model

import "math/big"

type Pitches []*Pitch

func (ps Pitches) MakePitchDurs(dur *big.Rat) []*PitchDur {
	n := len(ps)
	pds := make([]*PitchDur, n)

	for i := 0; i < n; i++ {
		pds[i] = NewPitchDur(ps[i], dur)
	}

	return pds
}
