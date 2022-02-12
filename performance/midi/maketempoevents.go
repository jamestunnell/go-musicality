package midi

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/performance/model"
	"github.com/rs/zerolog/log"
)

func MakeTempoEvents(
	fs *model.FlatScore,
	samplingDur rat.Rat,
	samplePeriod rat.Rat) ([]Event, error) {
	log.Debug().
		Str("sampling dur", samplingDur.String()).
		Str("sample period", samplePeriod.String()).
		Msg("collecting tempo events")

	tc, err := fs.TempoComputer()
	if err != nil {
		err = fmt.Errorf("failed to make tempo computer: %w", err)

		return []Event{}, err
	}

	offset := rat.Zero()
	bpm := tc.At(offset)
	events := []Event{
		NewTempoEvent(offset.Clone(), bpm),
	}

	offset.Accum(samplePeriod)

	for offset.Less(samplingDur) {
		newBPM := tc.At(offset)
		if newBPM != bpm {
			events = append(events, NewTempoEvent(offset.Clone(), newBPM))
			bpm = newBPM
		}

		offset.Accum(samplePeriod)
	}

	return events, nil
}
