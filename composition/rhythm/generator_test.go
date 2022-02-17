package rhythm_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/jamestunnell/go-musicality/common/function"
	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/composition/rhythm"
)

func TestGeneratorMakeMeasure(t *testing.T) {
	root := rhythm.NewNode(rat.New(1, 1))

	root.SubdivideRecursive(func(level int, n *rhythm.Node) (uint64, bool) {
		switch level {
		case 0:
			return 4, true
		case 1, 2, 3:
			return 2, true
		default:
			return 0, false
		}
	})

	g := rhythm.NewGenerator(root)

	for i := 0; i <= root.Depth(); i++ {
		t.Run(fmt.Sprintf("const depth %d", i), func(t *testing.T) {
			mDurs := g.Make(root.Duration(), function.NewConstantFunction(float64(i)))

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

			mDurs := g.Make(root.Duration(), function.NewRandomFunction(r))

			// t.Log(mDurs.Strings())

			assert.NotEmpty(t, mDurs)
			assert.True(t, mDurs.Sum().Equal(rat.New(1, 1)))
		})
	}
}

func TestGeneratorMakeDurDifferantThanRootDur(t *testing.T) {
	root := rhythm.NewNode(rat.New(1, 1))

	root.SubdivideRecursive(func(level int, n *rhythm.Node) (uint64, bool) {
		if n.Duration().LessEqual(rat.New(1, 16)) {
			return 0, false
		}

		return 2, true
	})

	f := function.NewConstantFunction(float64(2))
	makeDur := rat.New(7, 4)

	testGeneratorMake(t, root, f, makeDur)

	makeDur = rat.New(48, 50)

	testGeneratorMake(t, root, f, makeDur)
}

func testGeneratorMake(t *testing.T, root *rhythm.Node, f function.Function, makeDur rat.Rat) {
	g := rhythm.NewGenerator(root)
	mDurs := g.Make(makeDur, f)

	assert.NotEmpty(t, mDurs)
	assert.True(t, mDurs.Sum().Equal(makeDur))
}
