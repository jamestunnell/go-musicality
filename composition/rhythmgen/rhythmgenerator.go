package rhythmgen

import "github.com/jamestunnell/go-musicality/common/rat"

type RhythmGenerator interface {
	Reset()
	NextDur() rat.Rat
}
