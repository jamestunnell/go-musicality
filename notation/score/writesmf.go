package score

import (
	"fmt"
	"math/big"

	"github.com/rs/zerolog/log"
	"gitlab.com/gomidi/midi/smf"
	"gitlab.com/gomidi/midi/smf/smfwriter"
	"gitlab.com/gomidi/midi/writer"

	"github.com/jamestunnell/go-musicality/notation/meter"
)

const ticksPerQuarter = 960

func (s *Score) WriteSMF(fpath string) error {
	opts := []smfwriter.Option{
		smfwriter.NoRunningStatus(),
		smfwriter.TimeFormat(smf.MetricTicks(ticksPerQuarter)),
	}
	partNames := s.PartNames()
	numTracks := uint16(len(partNames))

	settings, err := s.MIDISettings()
	if err != nil {
		return fmt.Errorf("failed to get MIDI settings: %w", err)
	}

	for _, part := range partNames {
		if _, found := settings.PartChannels[part]; !found {
			return fmt.Errorf("part '%s' channel is missing from MIDI settings", part)
		}
	}

	write := func(wr *writer.SMF) error {
		for _, part := range partNames {
			log.Info().Str("name", part).Msg("processing part")

			if err := s.writePartSMF(wr, settings, part); err != nil {
				return err
			}
		}

		return nil
	}

	err = writer.WriteSMF(fpath, numTracks, write, opts...)
	if err != nil {
		return fmt.Errorf("failed to write SMF: %w", err)
	}

	return nil
}

func (s *Score) writePartSMF(wr *writer.SMF, settings *MIDISettings, part string) error {
	zero := big.NewRat(0, 0)

	writer.TrackSequenceName(wr, part)
	wr.SetChannel(settings.PartChannels[part])

	var met *meter.Meter

	for i, section := range s.Sections {
		log.Info().
			Str("name", section.Name).
			Int("index", i).
			Msg("processing section")

		for j, m := range section.Measures {
			if met == nil || !met.Equal(m.Meter) {
				met = m.Meter

				log.Info().
					Int("measure index", j).
					Str("meter", met.String()).
					Msg("setting meter")

				writer.Meter(wr, met.Numerator, met.Denominator)
			}

			notes, found := m.PartNotes[part]
			if found {
				remaining := m.Duration()
				for _, n := range notes {
					switch {
					case n.IsRest():
					case n.IsMonophonic():
						key, err := n.Pitches[0].MIDINote()
						if err != nil {
							return fmt.Errorf("failed to get MIDI key for note: %w", err)
						}
						writer.NoteOn(wr, key, 50)
						writer.Forward(wr, 0, uint32(n.Duration.Num().Uint64()), uint32(n.Duration.Denom().Uint64()))
						writer.NoteOff(wr, key)

						remaining.Sub(remaining, n.Duration)
					default:
						return fmt.Errorf("polyphonic notes not supported")
					}
				}

				if remaining.Cmp(zero) == 1 {
					writer.Forward(wr, 0, uint32(remaining.Num().Uint64()), uint32(remaining.Denom().Uint64()))
				}
			} else {
				writer.Forward(wr, 1, 0, 0)
			}
		}
	}

	writer.EndOfTrack(wr)

	return nil
}
