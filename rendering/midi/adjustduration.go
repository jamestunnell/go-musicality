package midi

import (
	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/note"
)

func AdjustDuration(dur rat.Rat, separation float64) rat.Rat {
	var adjust rat.Rat
	switch separation {
	case note.ControlMin:
		adjust = rat.New(1, 1)
	case note.ControlMax:
		adjust = rat.New(1, 8)
	case note.ControlNormal:
		adjust = rat.New(9, 16)
	default:
		remove := rat.New(7, 8).MulFloat64(separation)
		adjust = rat.New(1, 1).Sub(remove)
	}

	return dur.Mul(adjust)

}
