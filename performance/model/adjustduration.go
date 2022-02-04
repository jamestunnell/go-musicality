package sequence

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/validation"
)

func AdjustDuration(dur *big.Rat, separation float64) (*big.Rat, error) {
	err := validation.VerifyInRangeFloat(
		"separation", separation, note.SeparationMin, note.SeparationMax)
	if err != nil {
		return nil, err
	}

	var adjust *big.Rat
	switch separation {
	case note.SeparationMin:
		adjust = big.NewRat(1, 1)
	case note.SeparationMax:
		adjust = big.NewRat(1, 8)
	default:
		mul := (separation * 0.5) + 0.5
		remove := new(big.Rat).Mul(new(big.Rat).SetFloat64(mul), big.NewRat(7, 8))
		adjust = new(big.Rat).Sub(big.NewRat(1, 1), remove)
	}

	newDur := new(big.Rat).Mul(dur, adjust)

	return newDur, nil

}
