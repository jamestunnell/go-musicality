package model

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/note"
)

func AdjustDuration(dur *big.Rat, separation float64) *big.Rat {
	var adjust *big.Rat
	switch separation {
	case note.ControlMin:
		adjust = big.NewRat(1, 1)
	case note.ControlMax:
		adjust = big.NewRat(1, 8)
	default:
		remove := new(big.Rat).Mul(new(big.Rat).SetFloat64(separation), big.NewRat(7, 8))
		adjust = new(big.Rat).Sub(big.NewRat(1, 1), remove)
	}

	newDur := new(big.Rat).Mul(dur, adjust)

	return newDur

}
