package midi

import (
	"fmt"
	"math"
	"math/big"

	"github.com/rs/zerolog/log"
	"gitlab.com/gomidi/midi/smf"
	"gitlab.com/gomidi/midi/smf/smfwriter"
	"gitlab.com/gomidi/midi/writer"

	"github.com/jamestunnell/go-musicality/notation/score"
)

const ticksPerQuarter = 960

func WriteSMF(s *score.Score, fpath string) error {
	opts := []smfwriter.Option{
		smfwriter.NoRunningStatus(),
		smfwriter.TimeFormat(smf.MetricTicks(ticksPerQuarter)),
	}

	settings, err := Settings(s)
	if err != nil {
		return fmt.Errorf("failed to get MIDI settings: %w", err)
	}

	tracks, err := MakeTracks(s, settings)
	if err != nil {
		return fmt.Errorf("failed to make MIDI tracks: %w", err)
	}

	write := func(wr *writer.SMF) error {
		for _, track := range tracks {
			log.Info().Str("name", track.Name).Msg("writing track")

			writer.TrackSequenceName(wr, track.Name)
			wr.SetChannel(track.Channel)
			writer.ProgramChange(wr, track.Instrument)

			prev := big.NewRat(0, 1)

			for _, event := range track.Events {
				current := event.Offset

				if current.Cmp(prev) == 1 {
					diff := new(big.Rat).Sub(current, prev)

					flt, _ := diff.Float64()
					num := uint32(math.Round(math.MaxUint32 * flt))

					writer.Forward(wr, 0, num, math.MaxUint32)

					prev = current
				}

				err := event.Write(wr)
				if err != nil {
					return fmt.Errorf("failed to write event at %s: %w", event.Offset.String(), err)
				} else {
					offsetFlt, _ := current.Float64()
					log.Debug().Float64("offset", offsetFlt).Msg("wrote event")
				}
			}

			writer.EndOfTrack(wr)
		}

		return nil
	}

	err = writer.WriteSMF(fpath, uint16(len(tracks)), write, opts...)
	if err != nil {
		return fmt.Errorf("failed to write SMF: %w", err)
	}

	return nil
}

// func collectMeterEvents(fs *model.FlatScore) ([]*Event, error) {
// 	events := []*Event{}

// 	offset := big.NewRat(0, 1)

// 	var met *meter.Meter

// 	for _, section := range s.Sections {
// 		for _, m := range section.Measures {
// 			if met == nil || !met.Equal(m.Meter) {
// 				if m.Meter.Numerator >= 256 || m.Meter.Denominator >= 256 {
// 					return []*Event{}, fmt.Errorf("meter %s is not valid for MIDI", met.String())
// 				}

// 				metEvent := NewMeterEvent(offset, uint8(m.Meter.Numerator), uint8(m.Meter.Denominator))

// 				events = append(events, metEvent)

// 				met = m.Meter

// 				offset.Add(offset, m.Duration())
// 			}
// 		}
// 	}

// 	return events, nil
// }
