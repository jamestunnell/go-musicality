package sequence

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/note"
)

const (
	SeparationNone          = 0.0
	SeparationTenuto        = 0.2
	SeparationNormal        = 0.4
	SeparationPortato       = 0.6
	SeparationStaccato      = 0.8
	SeparationStaccatissimo = 1.0
)

func Separation(articulation string, slurring bool) float64 {
	sep := SeparationNormal

	if slurring {
		switch articulation {
		case note.Normal, note.Tenuto, note.Accent:
			sep = SeparationNone
		case note.Portato:
			sep = SeparationNormal
		case note.Marcato, note.Staccato:
			sep = SeparationPortato
		case note.Staccatissimo:
			sep = SeparationStaccato
		}
	} else {
		switch articulation {
		case note.Normal, note.Accent:
			sep = SeparationNormal
		case note.Tenuto:
			sep = SeparationTenuto
		case note.Marcato, note.Portato:
			sep = SeparationPortato
		case note.Staccato:
			sep = SeparationStaccato
		case note.Staccatissimo:
			sep = SeparationStaccatissimo
		}
	}

	return sep
}

func AdjustDuration(dur *big.Rat, separation float64) (*big.Rat, error) {
	if separation < 0.0 || separation > 1.0 {
		return nil, fmt.Errorf("separation %v is not in range [0,1]", separation)
	}

	var adjust *big.Rat
	switch separation {
	case 0.0:
		adjust = big.NewRat(1, 1)
	case 1.0:
		adjust = big.NewRat(1, 8)
	default:
		remove := new(big.Rat).Mul(new(big.Rat).SetFloat64(separation), big.NewRat(7, 8))
		adjust = new(big.Rat).Sub(big.NewRat(1, 1), remove)
	}

	newDur := new(big.Rat).Mul(dur, adjust)

	return newDur, nil

}
