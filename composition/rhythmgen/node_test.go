package rhythmgen_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/composition/rhythmgen"
)

func TestNodeSubdivideByZero(t *testing.T) {
	root := rhythmgen.NewNode(rat.New(1, 1))

	root.Subdivide(0)

	assert.Equal(t, 0, root.Depth())
}

func TestNodeSmallestDur(t *testing.T) {
	root := rhythmgen.NewNode(rat.New(1, 1))

	root.SubdivideRecursive(func(level int, n *rhythmgen.Node) (uint64, bool) {
		if n.Duration().Equal(rat.New(1, 32)) {
			return 0, false
		}

		return 2, true
	})

	assert.True(t, root.SmallestDur().Equal(rat.New(1, 32)))
}

func TestNodeVisit(t *testing.T) {
	root := rhythmgen.NewNode(rat.New(1, 1))

	root.SubdivideRecursive(func(level int, n *rhythmgen.Node) (uint64, bool) {
		if n.Duration().Equal(rat.New(1, 16)) {
			return 0, false
		}

		return 2, true
	})

	count := 0
	root.Visit(func(level int, n *rhythmgen.Node) bool {
		count++

		return true
	})

	assert.Equal(t, 16+8+4+2+1, count)
}

func TestNodeVisitTerminal(t *testing.T) {
	root := rhythmgen.NewNode(rat.New(1, 1))

	testNodeVisitTerminal(t, "root only - max level 0", root, 0, []string{"1/1"})
	testNodeVisitTerminal(t, "root only - max level 1", root, 1, []string{"1/1"})

	root.Subdivide(2)

	testNodeVisitTerminal(t, "split once - max level 0", root, 0, []string{"1/1"})
	testNodeVisitTerminal(t, "split once - max level 1", root, 1, []string{"1/2", "1/2"})

	root.Subs()[0].Subdivide(2)
	root.Subs()[1].Subdivide(3)

	testNodeVisitTerminal(t, "split twice - max level 0", root, 0, []string{"1/1"})
	testNodeVisitTerminal(t, "split twice - max level 1", root, 1, []string{"1/2", "1/2"})
	testNodeVisitTerminal(t, "split twice - max level 2", root, 2, []string{"1/4", "1/4", "1/6", "1/6", "1/6"})
}

func testNodeVisitTerminal(t *testing.T, name string, root *rhythmgen.Node, maxLevel int, expectedDurStrings []string) {
	t.Run(name, func(t *testing.T) {
		s := []string{}

		root.VisitTerminal(maxLevel, func(n *rhythmgen.Node) {
			s = append(s, n.Duration().String())
		})

		assert.Equal(t, expectedDurStrings, s)
	})
}
