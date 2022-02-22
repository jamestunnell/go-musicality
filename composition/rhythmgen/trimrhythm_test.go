package rhythmgen_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/composition/rhythmgen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTrimRhythm(t *testing.T) {
	totalDur := rat.New(1, 1)
	rhythmDurs := rat.Rats{
		rat.New(3, 4),
		rat.New(1, 2),
	}

	rhythmgen.TrimRhythm(totalDur, rhythmDurs)

	require.Len(t, rhythmDurs, 2)

	assert.True(t, rhythmDurs[0].Equal(rat.New(3, 4)))
	assert.True(t, rhythmDurs[1].Equal(rat.New(1, 4)))
}
