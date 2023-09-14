package rhythmgen_test

import (
	"math/big"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/composition/rhythmgen"
	"github.com/jamestunnell/go-musicality/composition/rhythmgen/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMakeRhythm(t *testing.T) {
	totalDur := big.NewRat(1, 1)
	g := mocks.NewMockRhythmGenerator(gomock.NewController(t))
	quarter := big.NewRat(1, 4)

	g.EXPECT().NextDur().Times(4).Return(quarter)

	rhythmDurs := rhythmgen.MakeRhythm(totalDur, g)

	assert.Len(t, rhythmDurs, 4)

	for _, dur := range rhythmDurs {
		assert.True(t, rat.IsEqual(dur, quarter))
	}
}
