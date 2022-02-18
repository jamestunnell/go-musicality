package rhythmgen

import "github.com/jamestunnell/go-musicality/common/rat"

type RhythmGenerator interface {
	MakeRhythm(dur rat.Rat) rat.Rats
}
