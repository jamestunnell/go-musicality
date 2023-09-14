package rhythmgen_test

import (
	"math/big"
	"testing"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/composition/rhythmgen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTrimRhythm(t *testing.T) {
	totalDur := big.NewRat(1, 1)
	rhythmDurs := rat.Rats{
		big.NewRat(3, 4),
		big.NewRat(1, 2),
	}

	rhythmgen.TrimRhythm(totalDur, rhythmDurs)

	require.Len(t, rhythmDurs, 2)

	assert.True(t, rat.IsEqual(rhythmDurs[0], big.NewRat(3, 4)))
	assert.True(t, rat.IsEqual(rhythmDurs[1], big.NewRat(1, 4)))
}
