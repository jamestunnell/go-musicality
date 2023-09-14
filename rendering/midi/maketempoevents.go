package midi

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/performance/flatscore"
	"github.com/rs/zerolog/log"
)

func MakeTempoEvents(
	fs *flatscore.FlatScore,
	samplingDur *big.Rat,
	samplePeriod *big.Rat) ([]Event, error) {
	log.Debug().
		Str("sampling dur", samplingDur.String()).
		Str("sample period", samplePeriod.String()).
		Msg("collecting tempo events")

	offset := rat.Zero()
	bpm := fs.TempoComputer.At(offset)
	events := []Event{
		NewTempoEvent(offset, bpm),
	}

	offset = rat.Add(offset, samplePeriod)

	for rat.IsLess(offset, samplingDur) {
		newBPM := fs.TempoComputer.At(offset)
		if newBPM != bpm {
			events = append(events, NewTempoEvent(offset, newBPM))
			bpm = newBPM
		}

		offset = rat.Add(offset, samplePeriod)
	}

	return events, nil
}
