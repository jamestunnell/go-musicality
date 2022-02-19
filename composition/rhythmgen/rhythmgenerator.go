package rhythmgen

import "github.com/jamestunnell/go-musicality/common/rat"

//go:generate mockgen -source=rhythmgenerator.go -destination=mocks/mockrhythmgenerator.go -package=mocks

type RhythmGenerator interface {
	Reset()
	NextDur() rat.Rat
}
