package midi

import (
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/performance/flatscore"
	"github.com/rs/zerolog/log"
)

func MakeTempoEvents(
	fs *flatscore.FlatScore,
	samplingDur rat.Rat,
	samplePeriod rat.Rat) ([]Event, error) {
	log.Debug().
		Str("sampling dur", samplingDur.String()).
		Str("sample period", samplePeriod.String()).
		Msg("collecting tempo events")

	offset := rat.Zero()
	bpm := fs.TempoComputer.At(offset)
	events := []Event{
		NewTempoEvent(offset.Clone(), bpm),
	}

	offset.Accum(samplePeriod)

	for offset.Less(samplingDur) {
		newBPM := fs.TempoComputer.At(offset)
		if newBPM != bpm {
			events = append(events, NewTempoEvent(offset.Clone(), newBPM))
			bpm = newBPM
		}

		offset.Accum(samplePeriod)
	}

	return events, nil
}
