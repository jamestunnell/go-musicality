package rhythmgen

import (
	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/meter"
)

func SubdivideMeasure(met *meter.Meter, smallestDur rat.Rat) *TreeNode {
	root := NewTreeNode(met.MeasureDuration())
	subdivideIfNextSubdurNotLessThanSmallest := func(level int, n *TreeNode) (uint64, bool) {
		subDur := n.Duration().Div(rat.FromUint64(2))
		if subDur.GreaterEqual(smallestDur) {
			return 2, true
		}

		return 0, false
	}

	root.SubdivideRecursive(func(level int, n *TreeNode) (uint64, bool) {
		switch level {
		case 0:
			return met.BeatsPerMeasure, true
		case 1:
			numer := met.BeatDuration.Num().Uint64()
			if numer > 1 {
				return numer, true
			}

			return subdivideIfNextSubdurNotLessThanSmallest(level, n)
		default:
			return subdivideIfNextSubdurNotLessThanSmallest(level, n)
		}
	})

	return root
}
