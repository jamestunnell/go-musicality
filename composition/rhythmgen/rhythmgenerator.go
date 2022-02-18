package rhythmgen

import "github.com/jamestunnell/go-musicality/common/rat"

type RhythmGenerator interface {
	NextDur() rat.Rat
}
