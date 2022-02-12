package model_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/performance/function"
	"github.com/jamestunnell/go-musicality/performance/model"
)

func TestTimeDelta(t *testing.T) {
	tc, err := model.NewComputer(120.0, change.Changes{})

	require.NoError(t, err)

	bdc, err := model.NewComputer(0.25, change.Changes{})

	require.NoError(t, err)

	xrange := function.NewRange(rat.Zero(), rat.New(1, 1))
	samplePeriod := rat.New(1, 16)

	dt, err := model.TimeDelta(tc, bdc, xrange, samplePeriod)

	require.NoError(t, err)

	assert.Equal(t, 2*time.Second, dt)
}

func TestTimeDeltaOneTempoChange(t *testing.T) {
	tc, err := model.NewComputer(120.0, change.Changes{
		change.NewImmediate(rat.New(1, 2), 150.0),
	})

	require.NoError(t, err)

	bdc, err := model.NewComputer(0.25, change.Changes{})

	require.NoError(t, err)

	xrange := function.NewRange(rat.Zero(), rat.New(1, 1))
	samplePeriod := rat.New(1, 16)

	dt, err := model.TimeDelta(tc, bdc, xrange, samplePeriod)

	require.NoError(t, err)

	assert.Equal(t, time.Duration(1.8*float64(time.Second)), dt)
}
