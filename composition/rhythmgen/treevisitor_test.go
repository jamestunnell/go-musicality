package rhythmgen_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/composition/rhythmgen"
)

func TestTreeVisitorVisitNext(t *testing.T) {
	root := rhythmgen.NewTreeNode(rat.New(1, 1))

	root.SubdivideRecursive(func(level int, n *rhythmgen.TreeNode) (uint64, bool) {
		switch level {
		case 0:
			return 4, true
		case 1, 2:
			return 2, true
		default:
			return 0, false
		}
	})

	require.True(t, root.SmallestDur().Equal(rat.New(1, 16)))

	onVisit := func(level int, n *rhythmgen.TreeNode) bool {
		return false
	}

	testTreeVisitorVisitNext(t, "root only (once)", root, onVisit, 1, []int{0}, []string{"1/1"})
	testTreeVisitorVisitNext(t, "root only (twice)", root, onVisit, 2, []int{0, 0}, []string{"1/1", "1/1"})

	onVisit = func(level int, n *rhythmgen.TreeNode) bool {
		return level == 0
	}

	testTreeVisitorVisitNext(t, "stop at first level (once)", root, onVisit, 5,
		[]int{0, 1, 1, 1, 1}, []string{"1/1", "1/4", "1/4", "1/4", "1/4"})

	testTreeVisitorVisitNext(t, "stop at first level (twice)", root, onVisit, 10,
		[]int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1},
		[]string{"1/1", "1/4", "1/4", "1/4", "1/4", "1/1", "1/4", "1/4", "1/4", "1/4"})

	onVisit = func(level int, n *rhythmgen.TreeNode) bool {
		return level < 2
	}

	testTreeVisitorVisitNext(t, "stop at second level (once)", root, onVisit, 13,
		[]int{0, 1, 2, 2, 1, 2, 2, 1, 2, 2, 1, 2, 2},
		[]string{"1/1", "1/4", "1/8", "1/8", "1/4", "1/8", "1/8", "1/4", "1/8", "1/8", "1/4", "1/8", "1/8"})
}

func TestTreeVisitorReset(t *testing.T) {
	root := rhythmgen.NewTreeNode(rat.New(1, 1))

	root.SubdivideRecursive(func(level int, n *rhythmgen.TreeNode) (uint64, bool) {
		if level == 2 {
			return 0, false
		}

		return 2, true
	})

	require.True(t, root.SmallestDur().Equal(rat.New(1, 4)))

	v := rhythmgen.NewTreeVisitor(root)

	v.VisitNext(func(level int, n *rhythmgen.TreeNode) bool {
		assert.Equal(t, 0, level)

		return true
	})

	v.VisitNext(func(level int, n *rhythmgen.TreeNode) bool {
		assert.Equal(t, 1, level)

		return true
	})

	v.VisitNext(func(level int, n *rhythmgen.TreeNode) bool {
		assert.Equal(t, 2, level)

		return true
	})

	v.Reset()

	v.VisitNext(func(level int, n *rhythmgen.TreeNode) bool {
		assert.Equal(t, 0, level)

		return true
	})
}

func testTreeVisitorVisitNext(
	t *testing.T,
	name string,
	root *rhythmgen.TreeNode,
	onVisit rhythmgen.OnVisitFunc,
	nIter int,
	expectedLevels []int,
	expectedDurStrings []string,
) {
	t.Run(name, func(t *testing.T) {
		require.Len(t, expectedDurStrings, nIter)
		require.Len(t, expectedLevels, nIter)

		v := rhythmgen.NewTreeVisitor(root)
		levels := make([]int, nIter)
		durStrings := make([]string, nIter)

		for i := 0; i < nIter; i++ {
			v.VisitNext(func(level int, n *rhythmgen.TreeNode) bool {
				levels[i] = level
				durStrings[i] = n.Duration().String()

				return onVisit(level, n)
			})
		}

		assert.Equal(t, expectedLevels, levels)
		assert.Equal(t, expectedDurStrings, durStrings)
	})
}
