package rhythmgen

type TreeVisitor struct {
	levels []int
	nodes  []*TreeNode
	index  int
}

func NewTreeVisitor(root *TreeNode) *TreeVisitor {
	nodes := []*TreeNode{}
	levels := []int{}

	root.Visit(func(level int, n *TreeNode) bool {
		nodes = append(nodes, n)
		levels = append(levels, level)

		return true
	})

	return &TreeVisitor{
		nodes:  nodes,
		levels: levels,
		index:  0,
	}
}

func (v *TreeVisitor) Reset() {
	v.index = 0
}

func (v *TreeVisitor) advance() {
	v.index++
	if v.index >= len(v.nodes) {
		v.index = 0
	}
}

func (v *TreeVisitor) VisitNext(onVisit OnVisitFunc) {
	level := v.levels[v.index]
	node := v.nodes[v.index]

	descend := onVisit(level, node)

	v.advance()

	if !descend {
		for v.levels[v.index] > level {
			v.advance()
		}
	}
}
