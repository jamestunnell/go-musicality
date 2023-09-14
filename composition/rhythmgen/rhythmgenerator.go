package rhythmgen

import "math/big"

//go:generate mockgen -source=rhythmgenerator.go -destination=mocks/mockrhythmgenerator.go -package=mocks

type RhythmGenerator interface {
	Reset()
	NextDur() *big.Rat
}
