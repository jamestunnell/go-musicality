package rhythmgen

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/common/rat"
)

func MakeRhythm(totalDur *big.Rat, g RhythmGenerator) rat.Rats {
	durs := rat.Rats{}

	for rat.IsLess(durs.Sum(), totalDur) {
		dur := g.NextDur()

		durs = append(durs, dur)
	}

	diff := rat.Sub(durs.Sum(), totalDur)
	if rat.IsPositive(diff) {
		last := durs[len(durs)-1]

		durs[len(durs)-1] = rat.Sub(last, diff)
	}

	return durs
}
