package midi

import (
	"fmt"
	"math"

	"github.com/rs/zerolog/log"
	"gitlab.com/gomidi/midi/smf"
	"gitlab.com/gomidi/midi/smf/smfwriter"
	"gitlab.com/gomidi/midi/writer"

	"github.com/jamestunnell/go-musicality/notation/rat"
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

			prev := rat.Zero()

			for _, event := range track.Events {
				current := event.Offset

				if current.Greater(prev) {
					diff := current.Sub(prev).Float64()
					num := uint32(math.Round(math.MaxUint32 * diff))

					writer.Forward(wr, 0, num, math.MaxUint32)

					prev = current
				}

				err := event.Write(wr)
				if err != nil {
					return fmt.Errorf("failed to write event at %s: %w", event.Offset.String(), err)
				} else {
					log.Debug().
						Float64("offset", current.Float64()).
						Str("summary", event.Writer.Summary()).
						Msg("wrote event")
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
