package rhythm_test

import (
	"fmt"
	"testing"

	"github.com/jamestunnell/go-musicality/composition/rhythm"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/rat"
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
			mDurs := g.MakeMeasure(rhythm.NewConstantSelector(i))

			// t.Log(mDurs.Strings())

			assert.NotEmpty(t, mDurs)
			assert.True(t, mDurs.Duration().Equal(rat.New(1, 1)))
		})

		t.Run(fmt.Sprintf("rand max depth %d", i), func(t *testing.T) {
			r := distuv.Binomial{
				N:   float64(i),
				P:   0.5,
				Src: rand.NewSource(1234),
			}

			mDurs := g.MakeMeasure(rhythm.NewRandomSelector(r))

			// t.Log(mDurs.Strings())

			assert.NotEmpty(t, mDurs)
			assert.True(t, mDurs.Duration().Equal(rat.New(1, 1)))
		})
	}
}

func TestGeneratorMakeMoreThanMeasure(t *testing.T) {
	met := meter.ThreeFour()
	smallestDur := rat.New(1, 32)
	g := rhythm.NewGenerator(met, smallestDur)
	dur := rat.New(7, 4)
	mDurs := g.Make(dur, rhythm.NewConstantSelector(g.Depth()/2))

	assert.NotEmpty(t, mDurs)
	assert.True(t, mDurs.Duration().Equal(dur))
}

func TestGeneratorMakeLessThanMeasure(t *testing.T) {
	met := meter.SixEight()
	smallestDur := rat.New(1, 64)
	g := rhythm.NewGenerator(met, smallestDur)
	dur := rat.New(48, 50)
	mDurs := g.Make(dur, rhythm.NewConstantSelector(g.Depth()/2))

	assert.NotEmpty(t, mDurs)
	assert.True(t, mDurs.Duration().Equal(dur))
}
