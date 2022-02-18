package rhythmgen

import "github.com/jamestunnell/go-musicality/common/rat"

type TreeNode struct {
	dur  rat.Rat
	subs []*TreeNode
}

type VisitFunc func(level int, n *TreeNode) bool
type SubdivideRecursiveFunc func(level int, n *TreeNode) (uint64, bool)

func NewTreeNode(dur rat.Rat) *TreeNode {
	return &TreeNode{
		dur:  dur,
		subs: []*TreeNode{},
	}
}

func (n *TreeNode) Depth() int {
	maxLevel := 0

	n.Visit(func(level int, n *TreeNode) bool {
		if level > maxLevel {
			maxLevel = level
		}

		return true
	})

	return maxLevel
}

func (n *TreeNode) SmallestDur() rat.Rat {
	smallest := n.dur

	n.Visit(func(level int, n *TreeNode) bool {
		if n.dur.Less(smallest) {
			smallest = n.dur
		}

		return true
	})

	return smallest
}

func (n *TreeNode) Subs() []*TreeNode {
	return n.subs
}

func (n *TreeNode) Duration() rat.Rat {
	return n.dur
}

func (n *TreeNode) Subdivide(divisor uint64) {
	if divisor == 0 {
		return
	}

	subdur := n.dur.Div(rat.FromUint64(divisor))
	subs := make([]*TreeNode, divisor)

	for i := uint64(0); i < divisor; i++ {
		subs[i] = NewTreeNode(subdur)
	}

	n.subs = subs
}

func (n *TreeNode) SubdivideRecursive(s SubdivideRecursiveFunc) {
	n.subdivideRecursive(0, s)
}

func (n *TreeNode) subdivideRecursive(level int, s SubdivideRecursiveFunc) {
	divisor, divide := s(level, n)
	if !divide {
		return
	}

	n.Subdivide(divisor)

	nextLevel := level + 1

	for _, sub := range n.subs {
		sub.subdivideRecursive(nextLevel, s)
	}
}

func (n *TreeNode) Terminal() bool {
	return len(n.subs) == 0
}

func (n *TreeNode) Visit(v VisitFunc) {
	n.visit(0, v)
}

func (n *TreeNode) VisitTerminal(maxLevel int, do func(*TreeNode)) {
	v := func(level int, node *TreeNode) bool {
		if level >= maxLevel || node.Terminal() {
			do(node)

			return false
		}

		return true
	}

	n.visit(0, v)
}

func (n *TreeNode) visit(level int, v VisitFunc) {
	if !v(level, n) {
		return
	}

	for _, sub := range n.subs {
		sub.visit(level+1, v)
	}
}
