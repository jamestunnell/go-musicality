package rhythmgen

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/common/rat"
)

func TrimRhythm(totalDur *big.Rat, rhythmDurs rat.Rats) {
	diff := rat.Sub(rhythmDurs.Sum(), totalDur)
	if rat.IsPositive(diff) {
		last := rhythmDurs[len(rhythmDurs)-1]

		rhythmDurs[len(rhythmDurs)-1] = rat.Sub(last, diff)
	}
}
