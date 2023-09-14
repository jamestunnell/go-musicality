package rhythmgen_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/composition/rhythmgen"
)

func TestTreeNodeSubdivideByZero(t *testing.T) {
	root := rhythmgen.NewTreeNode(big.NewRat(1, 1))

	root.Subdivide(0)

	assert.Equal(t, 0, root.Depth())
}

func TestTreeNodeSmallestDur(t *testing.T) {
	root := rhythmgen.NewTreeNode(big.NewRat(1, 1))

	root.SubdivideRecursive(func(level int, n *rhythmgen.TreeNode) (uint64, bool) {
		if rat.IsEqual(n.Duration(), big.NewRat(1, 32)) {
			return 0, false
		}

		return 2, true
	})

	assert.True(t, rat.IsEqual(root.SmallestDur(), big.NewRat(1, 32)))
}

func TestTreeNodeVisit(t *testing.T) {
	root := rhythmgen.NewTreeNode(big.NewRat(1, 1))

	root.SubdivideRecursive(func(level int, n *rhythmgen.TreeNode) (uint64, bool) {
		if rat.IsEqual(n.Duration(), big.NewRat(1, 16)) {
			return 0, false
		}

		return 2, true
	})

	count := 0
	root.Visit(func(level int, n *rhythmgen.TreeNode) bool {
		count++

		return true
	})

	assert.Equal(t, 16+8+4+2+1, count)
}

func TestTreeNodeVisitTerminal(t *testing.T) {
	root := rhythmgen.NewTreeNode(big.NewRat(1, 1))

	testTreeNodeVisitTerminal(t, "root only - max level 0", root, 0, []string{"1/1"})
	testTreeNodeVisitTerminal(t, "root only - max level 1", root, 1, []string{"1/1"})

	root.Subdivide(2)

	testTreeNodeVisitTerminal(t, "split once - max level 0", root, 0, []string{"1/1"})
	testTreeNodeVisitTerminal(t, "split once - max level 1", root, 1, []string{"1/2", "1/2"})

	root.Subs()[0].Subdivide(2)
	root.Subs()[1].Subdivide(3)

	testTreeNodeVisitTerminal(t, "split twice - max level 0", root, 0, []string{"1/1"})
	testTreeNodeVisitTerminal(t, "split twice - max level 1", root, 1, []string{"1/2", "1/2"})
	testTreeNodeVisitTerminal(t, "split twice - max level 2", root, 2, []string{"1/4", "1/4", "1/6", "1/6", "1/6"})
}

func testTreeNodeVisitTerminal(t *testing.T, name string, root *rhythmgen.TreeNode, maxLevel int, expectedDurStrings []string) {
	t.Run(name, func(t *testing.T) {
		s := []string{}

		root.VisitTerminal(maxLevel, func(n *rhythmgen.TreeNode) {
			s = append(s, n.Duration().String())
		})

		assert.Equal(t, expectedDurStrings, s)
	})
}
