package rhythm_test

import (
	"fmt"
	"testing"

	"github.com/jamestunnell/go-musicality/composition/rhythm"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/performance/function"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

func TestGeneratorMakeMeasure(t *testing.T) {
	met := meter.FourFour()
	smallestDur := rat.New(1, 16)
	g := rhythm.NewGenerator(met, smallestDur)

	assert.True(t, g.SmallestDur().Equal(smallestDur))

	for i := 0; i <= g.Depth(); i++ {
		t.Run(fmt.Sprintf("const depth %d", i), func(t *testing.T) {
			mDurs := g.MakeMeasure(function.NewConstantFunction(float64(i)))

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

			mDurs := g.MakeMeasure(function.NewRandomFunction(r))

			// t.Log(mDurs.Strings())

			assert.NotEmpty(t, mDurs)
			assert.True(t, mDurs.Sum().Equal(rat.New(1, 1)))
		})
	}
}

func TestGeneratorMakeMoreThanMeasure(t *testing.T) {
	f := function.NewConstantFunction(float64(2))
	makeDur := rat.New(7, 4)

	testGeneratorMake(t, meter.ThreeFour(), rat.New(1, 32), f, makeDur)
}

func TestGeneratorMakeLessThanMeasure(t *testing.T) {
	f := function.NewConstantFunction(float64(2))
	makeDur := rat.New(48, 50)

	testGeneratorMake(t, meter.SixEight(), rat.New(1, 64), f, makeDur)
}

func testGeneratorMake(t *testing.T, met *meter.Meter, smallestDur rat.Rat, f function.Function, makeDur rat.Rat) {
	g := rhythm.NewGenerator(met, smallestDur)
	mDurs := g.Make(makeDur, f)

	assert.NotEmpty(t, mDurs)
	assert.True(t, mDurs.Sum().Equal(makeDur))
}
