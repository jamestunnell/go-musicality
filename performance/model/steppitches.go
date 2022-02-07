package model

import (
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

func MakeSteps(dur rat.Rat, start, end *pitch.Pitch, centsPerStep int) []*PitchDur {
	ps := StepPitches(start, end, centsPerStep)
	nSteps := rat.FromInt64(int64(len(ps)))
	subDur := dur.Div(nSteps)

	return ps.MakePitchDurs(subDur)
}

func StepPitches(startPitch, endPitch *pitch.Pitch, centsPerStep int) Pitches {
	start := NewPitch(startPitch, 0).TotalCent()
	end := NewPitch(endPitch, 0).TotalCent()

	diff := end - start
	if diff < 0 {
		centsPerStep = -centsPerStep
	}

	nSteps := diff / centsPerStep
	diffCurrent := 0
	pitches := make(Pitches, nSteps)

	for i := 0; i < nSteps; i++ {
		pitches[i] = NewPitch(startPitch, diffCurrent)

		diffCurrent += centsPerStep
	}

	return pitches
}
