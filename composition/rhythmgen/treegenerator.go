package rhythmgen

import (
	"github.com/jamestunnell/go-musicality/common/function"
	"github.com/jamestunnell/go-musicality/common/rat"
)

type TreeGenerator struct {
	root *TreeNode
}

func NewTreeGenerator(root *TreeNode) *TreeGenerator {
	return &TreeGenerator{root: root}
}

func (g *TreeGenerator) Make(dur rat.Rat, maxLevelFunction function.Function) rat.Rats {
	durs := rat.Rats{}
	x := rat.Zero()
	maxLevel := int(maxLevelFunction.At(x))
	done := false

	for durs.Sum().Less(dur) {
		g.root.Visit(func(level int, n *TreeNode) bool {
			if done {
				return false
			}

			if level >= maxLevel || n.Terminal() {
				durs = append(durs, n.Duration())
				x = x.Add(n.Duration())
				maxLevel = int(maxLevelFunction.At(x))
				done = x.GreaterEqual(dur)

				return false
			}

			return true
		})
	}

	diff := durs.Sum().Sub(dur)
	if diff.Positive() {
		last := durs[len(durs)-1]

		durs[len(durs)-1] = last.Sub(diff)
	}

	return durs
}
