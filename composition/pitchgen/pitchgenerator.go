package pitchgen

import "github.com/jamestunnell/go-musicality/notation/pitch"

type PitchGenerator interface {
	MakePitches(n int) pitch.Pitches
}
