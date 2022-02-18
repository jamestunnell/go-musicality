package pitchgen

import "github.com/jamestunnell/go-musicality/notation/pitch"

func MakePitches(n int, g PitchGenerator) pitch.Pitches {
	pitches := make(pitch.Pitches, n)

	for i := 0; i < n; i++ {
		pitches[i] = g.NextPitch()
	}

	return pitches
}
