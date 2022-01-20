package midi

import (
	"fmt"
	"math"
	"math/big"

	"github.com/rs/zerolog/log"
	"gitlab.com/gomidi/midi/smf"
	"gitlab.com/gomidi/midi/smf/smfwriter"
	"gitlab.com/gomidi/midi/writer"

	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/jamestunnell/go-musicality/performance/sequence"
)

const ticksPerQuarter = 960

func WriteSMF(s *score.Score, fpath string) error {
	opts := []smfwriter.Option{
		smfwriter.NoRunningStatus(),
		smfwriter.TimeFormat(smf.MetricTicks(ticksPerQuarter)),
	}
	partNames := s.PartNames()
	numTracks := uint16(len(partNames))

	settings, err := Settings(s)
	if err != nil {
		return fmt.Errorf("failed to get MIDI settings: %w", err)
	}

	for _, part := range partNames {
		if _, found := settings.PartChannels[part]; !found {
			return fmt.Errorf("part '%s' channel is missing from MIDI settings", part)
		}
	}

	write := func(wr *writer.SMF) error {
		metEvents, err := collectMeterEvents(s)
		if err != nil {
			return fmt.Errorf("failed to collect meter events: %w", err)
		}

		for _, part := range partNames {
			log.Info().Str("name", part).Msg("processing part")

			noteEvents, err := collectNoteEvents(s, part)
			if err != nil {
				return fmt.Errorf("failed to collect notes for part '%s': %w", part, err)
			}

			writer.TrackSequenceName(wr, part)
			wr.SetChannel(settings.PartChannels[part])

			events := append(metEvents, noteEvents...)

			if len(events) == 0 {
				break
			}

			SortEvents(events)

			prev := big.NewRat(0, 1)

			for _, event := range events {
				current := event.Offset
				diff := new(big.Rat).Sub(current, prev)

				if diff.Cmp(big.NewRat(0, 1)) == 1 {
					writer.Forward(wr, 0, uint32(diff.Num().Uint64()), uint32(diff.Denom().Uint64()))
				}

				if err := event.Write(wr); err != nil {
					return fmt.Errorf("failed to write event at %s: %w", event.Offset.String(), err)
				}
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

func collectMeterEvents(s *score.Score) ([]*Event, error) {
	events := []*Event{}

	offset := big.NewRat(0, 1)

	var met *meter.Meter

	for _, section := range s.Sections {
		for _, m := range section.Measures {
			if met == nil || !met.Equal(m.Meter) {
				if m.Meter.Numerator >= 256 || m.Meter.Denominator >= 256 {
					return []*Event{}, fmt.Errorf("meter %s is not valid for MIDI", met.String())
				}

				metEvent := NewMeterEvent(offset, uint8(m.Meter.Numerator), uint8(m.Meter.Denominator))

				events = append(events, metEvent)

				met = m.Meter

				offset.Add(offset, m.Duration())
			}
		}
	}

	return events, nil
}

func collectNoteEvents(s *score.Score, part string) ([]*Event, error) {
	events := []*Event{}

	notes := s.PartNotes(part)
	seqs := sequence.Extract(notes)

	for _, seq := range seqs {
		offsets := seq.Offsets()

		for i, elem := range seq.Elements {
			p := elem.Pitch
			key, err := Key(p)
			if err != nil {
				err = fmt.Errorf("failed to get MIDI note for pitch '%s': %w", p.String(), err)

				return []*Event{}, err
			}

			vel, err := Velocity(elem.Attack)
			if err != nil {
				err = fmt.Errorf("failed to get MIDI velocity: %w", err)

				return []*Event{}, err
			}

			NewNoteOnEvent(offsets[i], key, vel)

			lastElem := i == (len(seq.Elements) - 1)

			if lastElem {
				newDur, err := sequence.AdjustDuration(elem.Duration, seq.Separation)
				if err != nil {
					err = fmt.Errorf("failed to adjust duration: %w", err)

					return []*Event{}, err
				}

				endOffset := new(big.Rat).Add(offsets[i], newDur)

				NewNoteOffEvent(endOffset, key)
			} else {
				NewNoteOffEvent(offsets[i+1], key)
			}
		}
	}

	return events, nil
}

// func writePartSMF(s *score.Score, wr *writer.SMF, settings *MIDISettings, part string) error {
// 	zero := big.NewRat(0, 0)

// 	writer.TrackSequenceName(wr, part)
// 	wr.SetChannel(settings.PartChannels[part])

// 	var met *meter.Meter

// 	for i, section := range s.Sections {
// 		log.Info().
// 			Str("name", section.Name).
// 			Int("index", i).
// 			Msg("processing section")

// 		for j, m := range section.Measures {
// 			if met == nil || !met.Equal(m.Meter) {
// 				met = m.Meter

// 				log.Info().
// 					Int("measure index", j).
// 					Str("meter", met.String()).
// 					Msg("setting meter")

// 				writer.Meter(wr, met.Numerator, met.Denominator)
// 			}

// 			notes, found := m.PartNotes[part]
// 			if found {
// 				remaining := m.Duration()
// 				for _, n := range notes {
// 					switch {
// 					case n.IsRest():
// 					case n.IsMonophonic():
// 						key, err := n.Pitches[0].MIDINote()
// 						if err != nil {
// 							return fmt.Errorf("failed to get MIDI key for note: %w", err)
// 						}
// 						writer.NoteOn(wr, key, 50)
// 						writer.Forward(wr, 0, uint32(n.Duration.Num().Uint64()), uint32(n.Duration.Denom().Uint64()))
// 						writer.NoteOff(wr, key)

// 						remaining.Sub(remaining, n.Duration)
// 					default:
// 						return fmt.Errorf("polyphonic notes not supported")
// 					}
// 				}

// 				if remaining.Cmp(zero) == 1 {
// 					writer.Forward(wr, 0, uint32(remaining.Num().Uint64()), uint32(remaining.Denom().Uint64()))
// 				}
// 			} else {
// 				writer.Forward(wr, 1, 0, 0)
// 			}
// 		}
// 	}

// 	writer.EndOfTrack(wr)

// 	return nil
// }

// Key converts the pitch to a MIDI note number.
// Returns a non-nil error if the pitch is not in the range [C-1, G9].
func Key(p *pitch.Pitch) (uint8, error) {
	const (
		// minTotalSemitone is the total semitone value of MIDI note 0 (octave below C0)
		minTotalSemitone = -12
		// maxTotalSemitone is the total semitone value of MIDI note 127 (G9)
		maxTotalSemitone = 115
	)

	totalSemitone := p.TotalSemitone()

	if totalSemitone < minTotalSemitone || totalSemitone > maxTotalSemitone {
		return 0, fmt.Errorf("pitch %s is outside of MIDI note number range", p.String())
	}

	return uint8(totalSemitone + 12), nil
}

// Velocity converts the attack to a MIDI velocity.
// Returns a non-nil error if the attack is not in the range [0.0, 1.0]
func Velocity(attack float64) (uint8, error) {
	if attack < 0.0 || attack > 1.0 {
		return 0, fmt.Errorf("attack '%v' not in range [0.0, 1.0]", attack)
	}

	vel := uint8(math.Round(attack * 127))

	return vel, nil
}
