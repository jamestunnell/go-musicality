package rhythmgen

import (
	"github.com/jamestunnell/go-musicality/common/function"
	"github.com/jamestunnell/go-musicality/common/rat"
)

type TreeGenerator struct {
	root     *TreeNode
	maxLevel function.Function
}

func NewTreeGenerator(root *TreeNode, maxLevel function.Function) *TreeGenerator {
	return &TreeGenerator{root: root, maxLevel: maxLevel}
}

func (g *TreeGenerator) MakeRhythm(dur rat.Rat) rat.Rats {
	durs := rat.Rats{}
	x := rat.Zero()
	maxLevel := int(g.maxLevel.At(x))
	done := false

	for durs.Sum().Less(dur) {
		g.root.Visit(func(level int, n *TreeNode) bool {
			if done {
				return false
			}

			if level >= maxLevel || n.Terminal() {
				durs = append(durs, n.Duration())
				x = x.Add(n.Duration())
				maxLevel = int(g.maxLevel.At(x))
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
