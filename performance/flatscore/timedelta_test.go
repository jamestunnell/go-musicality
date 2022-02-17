package flatscore_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/common/function"
	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/performance/computer"
	"github.com/jamestunnell/go-musicality/performance/flatscore"
)

func TestTimeDelta(t *testing.T) {
	tc, err := computer.New(120.0, change.Changes{})

	require.NoError(t, err)

	bdc, err := computer.New(0.25, change.Changes{})

	require.NoError(t, err)

	xrange := function.NewRange(rat.Zero(), rat.New(1, 1))
	samplePeriod := rat.New(1, 16)

	dt, err := flatscore.TimeDelta(tc, bdc, xrange, samplePeriod)

	require.NoError(t, err)

	assert.Equal(t, 2*time.Second, dt)
}

func TestTimeDeltaOneTempoChange(t *testing.T) {
	tc, err := computer.New(120.0, change.Changes{
		change.NewImmediate(rat.New(1, 2), 150.0),
	})

	require.NoError(t, err)

	bdc, err := computer.New(0.25, change.Changes{})

	require.NoError(t, err)

	xrange := function.NewRange(rat.Zero(), rat.New(1, 1))
	samplePeriod := rat.New(1, 16)

	dt, err := flatscore.TimeDelta(tc, bdc, xrange, samplePeriod)

	require.NoError(t, err)

	assert.Equal(t, time.Duration(1.8*float64(time.Second)), dt)
}
