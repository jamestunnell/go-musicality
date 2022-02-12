package model

import (
	"fmt"
	"time"

	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/performance/function"
)

// TimeDelta takes two note offsets, and uses the tempo and beat duration computers to
// determine how much time passes between them.
func TimeDelta(
	tempoComp, beatDurComp *Computer,
	xrange function.Range,
	samplePeriod rat.Rat) (time.Duration, error) {
	bpms, err := function.Sample(tempoComp, xrange, samplePeriod)
	if err != nil {
		return 0, fmt.Errorf("failed to sample tempos: %w", err)
	}

	bdurs, err := function.Sample(beatDurComp, xrange, samplePeriod)
	if err != nil {
		return 0, fmt.Errorf("failed to sample beat durs: %w", err)
	}

	deltaSec := 0.0
	samplePeriodFlt := samplePeriod.Float64()

	for i := 0; i < (len(bpms) - 1); i++ {
		deltaSec += (60 * samplePeriodFlt) / (bdurs[i] * bpms[i])
	}

	timeDelta := time.Duration(deltaSec * float64(time.Second))

	return timeDelta, nil
}
