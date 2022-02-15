package mononote

import (
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/performance/centpitch"
	"github.com/jamestunnell/go-musicality/performance/pitchdur"
)

func MakeSteps(dur rat.Rat, start, end *pitch.Pitch, centsPerStep int) []*pitchdur.PitchDur {
	ps := MakeStepPitches(start, end, centsPerStep)
	nSteps := rat.FromInt64(int64(len(ps)))
	subDur := dur.Div(nSteps)

	return MakePitchDurs(subDur, ps)
}

func MakeStepPitches(startPitch, endPitch *pitch.Pitch, centsPerStep int) []*centpitch.CentPitch {
	start := centpitch.New(startPitch, 0).TotalCent()
	end := centpitch.New(endPitch, 0).TotalCent()

	diff := end - start
	if diff < 0 {
		centsPerStep = -centsPerStep
	}

	nSteps := diff / centsPerStep
	diffCurrent := 0
	pitches := make([]*centpitch.CentPitch, nSteps)

	for i := 0; i < nSteps; i++ {
		pitches[i] = centpitch.New(startPitch, diffCurrent)

		diffCurrent += centsPerStep
	}

	return pitches
}

func MakePitchDurs(dur rat.Rat, pitches []*centpitch.CentPitch) []*pitchdur.PitchDur {
	n := len(pitches)
	pds := make([]*pitchdur.PitchDur, n)

	for i := 0; i < n; i++ {
		pds[i] = pitchdur.New(pitches[i], dur)
	}

	return pds
}
