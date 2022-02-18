package rhythmgen

import "github.com/jamestunnell/go-musicality/common/rat"

type Node struct {
	dur  rat.Rat
	subs []*Node
}

type VisitFunc func(level int, n *Node) bool
type SubdivideRecursiveFunc func(level int, n *Node) (uint64, bool)

func NewNode(dur rat.Rat) *Node {
	return &Node{
		dur:  dur,
		subs: []*Node{},
	}
}

func (n *Node) Depth() int {
	maxLevel := 0

	n.Visit(func(level int, n *Node) bool {
		if level > maxLevel {
			maxLevel = level
		}

		return true
	})

	return maxLevel
}

func (n *Node) SmallestDur() rat.Rat {
	smallest := n.dur

	n.Visit(func(level int, n *Node) bool {
		if n.dur.Less(smallest) {
			smallest = n.dur
		}

		return true
	})

	return smallest
}

func (n *Node) Subs() []*Node {
	return n.subs
}

func (n *Node) Duration() rat.Rat {
	return n.dur
}

func (n *Node) Subdivide(divisor uint64) {
	if divisor == 0 {
		return
	}

	subdur := n.dur.Div(rat.FromUint64(divisor))
	subs := make([]*Node, divisor)

	for i := uint64(0); i < divisor; i++ {
		subs[i] = NewNode(subdur)
	}

	n.subs = subs
}

func (n *Node) SubdivideRecursive(s SubdivideRecursiveFunc) {
	n.subdivideRecursive(0, s)
}

func (n *Node) subdivideRecursive(level int, s SubdivideRecursiveFunc) {
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

func (n *Node) Terminal() bool {
	return len(n.subs) == 0
}

func (n *Node) Visit(v VisitFunc) {
	n.visit(0, v)
}

func (n *Node) VisitTerminal(maxLevel int, do func(*Node)) {
	v := func(level int, node *Node) bool {
		if level >= maxLevel || node.Terminal() {
			do(node)

			return false
		}

		return true
	}

	n.visit(0, v)
}

func (n *Node) visit(level int, v VisitFunc) {
	if !v(level, n) {
		return
	}

	for _, sub := range n.subs {
		sub.visit(level+1, v)
	}
}
