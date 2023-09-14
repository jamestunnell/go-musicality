package midi

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/note"
)

func AdjustDuration(dur *big.Rat, separation float64) *big.Rat {
	var adjust *big.Rat
	switch separation {
	case note.ControlMin:
		adjust = big.NewRat(1, 1)
	case note.ControlMax:
		adjust = big.NewRat(1, 8)
	case note.ControlNormal:
		adjust = big.NewRat(9, 16)
	default:
		remove := rat.Mul(big.NewRat(7, 8), rat.FromFloat64(separation))
		adjust = rat.Sub(big.NewRat(1, 1), remove)
	}

	return rat.Mul(dur, adjust)

}
