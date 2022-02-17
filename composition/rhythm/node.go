package rhythm

import "github.com/jamestunnell/go-musicality/notation/rat"

type Node struct {
	elem *Element
	subs []*Node
}

type VisitFunc func(level int, n *Node) bool

func NewNode(elem *Element) *Node {
	return &Node{
		elem: elem,
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
	smallest := n.elem.Duration

	n.Visit(func(level int, n *Node) bool {
		if n.elem.Duration.Less(smallest) {
			smallest = n.elem.Duration
		}

		return true
	})

	return smallest
}

func (n *Node) Subs() []*Node {
	return n.subs
}

func (n *Node) Element() *Element {
	return n.elem
}

func (n *Node) Subdivide(divisor uint64) {
	subs := make([]*Node, divisor)
	for i, e := range n.elem.Divide(divisor) {
		subs[i] = NewNode(e)
	}

	n.subs = subs
}

func (n *Node) SubdivideN(divisor, nTimes uint64) {
	if nTimes > 0 {
		n.Subdivide(divisor)

		for _, sub := range n.subs {
			sub.SubdivideN(divisor, nTimes-1)
		}
	}
}

func (n *Node) SubdivideUntil(divisor uint64, check func(*Node) bool) {
	if check(n) {
		n.Subdivide(divisor)

		for _, sub := range n.subs {
			sub.SubdivideUntil(divisor, check)
		}
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
