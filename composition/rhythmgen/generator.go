package rhythmgen

import (
	"github.com/jamestunnell/go-musicality/common/function"
	"github.com/jamestunnell/go-musicality/common/rat"
)

type Generator struct {
	root *Node
}

func NewGenerator(root *Node) *Generator {
	return &Generator{root: root}
}

func (g *Generator) Make(dur rat.Rat, maxLevelFunction function.Function) rat.Rats {
	durs := rat.Rats{}
	x := rat.Zero()
	maxLevel := int(maxLevelFunction.At(x))
	done := false

	for durs.Sum().Less(dur) {
		g.root.Visit(func(level int, n *Node) bool {
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
