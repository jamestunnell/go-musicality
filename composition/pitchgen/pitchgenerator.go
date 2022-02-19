package pitchgen

import "github.com/jamestunnell/go-musicality/notation/pitch"

type PitchGenerator interface {
	Reset()
	NextPitch() *pitch.Pitch
}
