package rhythmgen_test

import (
	"fmt"
	"testing"

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

	for i := 0; i <= root.Depth(); i++ {
		t.Run(fmt.Sprintf("const depth %d", i), func(t *testing.T) {
			maxLevel := function.NewConstantFunction(float64(i))
			g := rhythmgen.NewTreeGenerator(root, maxLevel)

			mDurs := rhythmgen.MakeRhythm(root.Duration(), g)

			// t.Log(mDurs.Strings())

			assert.NotEmpty(t, mDurs)
			assert.True(t, mDurs.Sum().Equal(rat.New(1, 1)))
		})

		t.Run(fmt.Sprintf("rand max depth %d", i), func(t *testing.T) {
			r := distuv.Binomial{
				N:   float64(i),
				P:   0.5,
				Src: rand.NewSource(1234),
			}
			maxLevel := function.NewRandomFunction(r)
			g := rhythmgen.NewTreeGenerator(root, maxLevel)

			mDurs := rhythmgen.MakeRhythm(root.Duration(), g)

			// t.Log(mDurs.Strings())

			assert.NotEmpty(t, mDurs)
			assert.True(t, mDurs.Sum().Equal(rat.New(1, 1)))
		})
	}
}

func TestTreeGeneratorMakeDurDifferantThanRootDur(t *testing.T) {
	root := rhythmgen.NewTreeNode(rat.New(1, 1))

	root.SubdivideRecursive(func(level int, n *rhythmgen.TreeNode) (uint64, bool) {
		if n.Duration().LessEqual(rat.New(1, 16)) {
			return 0, false
		}

		return 2, true
	})

	f := function.NewConstantFunction(float64(2))
	makeDur := rat.New(7, 4)

	testTreeGeneratorMake(t, root, f, makeDur)

	makeDur = rat.New(48, 50)

	testTreeGeneratorMake(t, root, f, makeDur)
}

func testTreeGeneratorMake(t *testing.T, root *rhythmgen.TreeNode, f function.Function, makeDur rat.Rat) {
	g := rhythmgen.NewTreeGenerator(root, f)
	mDurs := rhythmgen.MakeRhythm(makeDur, g)

	assert.NotEmpty(t, mDurs)
	assert.True(t, mDurs.Sum().Equal(makeDur))
}
