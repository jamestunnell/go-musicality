package rhythmgen

import (
	"github.com/jamestunnell/go-musicality/common/function"
	"github.com/jamestunnell/go-musicality/common/rat"
)

type TreeGenerator struct {
	maxLevel        int
	maxLevelFunc    function.Function
	visitor         *TreeVisitor
	latestDur       rat.Rat
	durSoFar        rat.Rat
	reachedTerminal bool
}

func NewTreeGenerator(root *TreeNode, maxLevelFunc function.Function) *TreeGenerator {
	zero := rat.Zero()
	g := &TreeGenerator{
		reachedTerminal: false,
		visitor:         NewTreeVisitor(root),
		maxLevelFunc:    maxLevelFunc,
		maxLevel:        int(maxLevelFunc.At(zero)),
		latestDur:       zero,
		durSoFar:        zero,
	}

	return g
}

func (g *TreeGenerator) NextDur() rat.Rat {
	g.reachedTerminal = false
	for !g.reachedTerminal {
		g.visitor.VisitNext(g.onVisit)
	}

	return g.latestDur
}

func (g *TreeGenerator) onVisit(level int, n *TreeNode) bool {
	if level >= g.maxLevel || n.Terminal() {
		g.reachedTerminal = true
		g.latestDur = n.Duration()
		g.durSoFar = g.durSoFar.Add(g.latestDur)
		g.maxLevel = int(g.maxLevelFunc.At(g.durSoFar))

		return false
	}

	return true
}
