package rhythmgen_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/jamestunnell/go-musicality/common/function"
	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/composition/rhythmgen"
)

func TestTreeGeneratorMakeMeasure(t *testing.T) {
	root := rhythmgen.NewTreeNode(rat.New(1, 1))

	root.SubdivideRecursive(func(level int, n *rhythmgen.TreeNode) (uint64, bool) {
		switch level {
		case 0:
			return 4, true
		case 1, 2, 3:
			return 2, true
		default:
			return 0, false
		}
	})

	testTreeGeneratorConst(t, root, 0, "1/1")
	testTreeGeneratorConst(t, root, 1, "1/4", "1/4", "1/4", "1/4")
	testTreeGeneratorConst(t, root, 2, "1/8", "1/8", "1/8", "1/8", "1/8", "1/8", "1/8", "1/8")
	testTreeGeneratorRand(t, root, 0)
	testTreeGeneratorRand(t, root, 1)
	testTreeGeneratorRand(t, root, 2)
	testTreeGeneratorRand(t, root, root.Depth())
	testTreeGeneratorReset(t, root)
}

func testTreeGeneratorConst(
	t *testing.T, root *rhythmgen.TreeNode, maxLevel int, expectedDurStrings ...string) {
	t.Run(fmt.Sprintf("const level %d", maxLevel), func(t *testing.T) {
		f := function.NewConstantFunction(float64(maxLevel))
		g := rhythmgen.NewTreeGenerator(root, f)
		smallest := root.SmallestDur()

		for _, expectedDurStr := range expectedDurStrings {
			dur := g.NextDur()

			assert.True(t, dur.Positive())
			assert.True(t, dur.GreaterEqual(smallest))
			assert.Equal(t, expectedDurStr, dur.String())
		}
	})
}

func testTreeGeneratorRand(t *testing.T, root *rhythmgen.TreeNode, maxLevel int) {
	t.Run(fmt.Sprintf("rand level %d", maxLevel), func(t *testing.T) {
		dist := distuv.Binomial{
			N:   float64(maxLevel),
			P:   0.5,
			Src: rand.NewSource(uint64(time.Now().Unix())),
		}
		f := function.NewRandomFunction(dist)
		g := rhythmgen.NewTreeGenerator(root, f)
		smallest := root.SmallestDur()

		for i := 0; i < 10; i++ {
			dur := g.NextDur()

			assert.True(t, dur.Positive())
			assert.True(t, dur.GreaterEqual(smallest))
		}
	})
}

func testTreeGeneratorReset(t *testing.T, root *rhythmgen.TreeNode) {
	t.Run("reset", func(t *testing.T) {
		maxLevel := function.NewConstantFunction(2)
		g := rhythmgen.NewTreeGenerator(root, maxLevel)

		mDurs1 := rhythmgen.MakeRhythm(root.Duration(), g)

		g.Reset()

		mDurs2 := rhythmgen.MakeRhythm(root.Duration(), g)

		assert.NotEmpty(t, mDurs1)
		assert.True(t, mDurs2.Equal(mDurs1))
	})
}
