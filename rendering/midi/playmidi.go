package midi

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/rs/zerolog/log"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
	"gitlab.com/gomidi/rtmididrv"

	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/jamestunnell/go-musicality/performance/flatscore"
	"github.com/jamestunnell/go-musicality/performance/function"
)

type playStep struct {
	Offset      rat.Rat
	TrackEvents map[int]NoteEvents
}

var errNoMIDIOutputs = errors.New("no MIDI output ports")

func PlayMIDI(s *score.Score) error {
	settings, err := Settings(s)
	if err != nil {
		return fmt.Errorf("failed to get MIDI settings: %w", err)
	}

	fs, err := flatscore.Convert(s)
	if err != nil {
		return fmt.Errorf("failed to convert to flat score: %w", err)
	}

	tracks, err := MakeTracks(fs, settings)
	if err != nil {
		return fmt.Errorf("failed to make MIDI tracks: %w", err)
	}

	if len(tracks) == 0 {
		return fmt.Errorf("score produced no MIDI tracks")
	}

	drv, err := rtmididrv.New()
	if err != nil {
		return fmt.Errorf("failed to open MIDI driver: %w", err)
	}

	outs, err := drv.Outs()
	if err != nil {
		return fmt.Errorf("failed to get MIDI outputs: %w", err)
	}

	if len(outs) == 0 {
		return errNoMIDIOutputs
	}

	out := outs[0]

	if err = out.Open(); err != nil {
		return fmt.Errorf("failed to open MIDI output port %s: %w", out.String(), err)
	}

	// make sure to close all open ports at the end
	defer drv.Close()

	return playMIDI(settings, out, tracks, fs)
}

func playMIDI(settings *MIDISettings, out midi.Out, tracks []*Track, fs *flatscore.FlatScore) error {
	log.Info().Str("name", out.String()).Msg("opened MIDI output port")

	offset := rat.Zero()
	wr := writer.New(out)
	steps := makePlaySteps(tracks)

	for _, step := range steps {
		log.Debug().Str("offset", step.Offset.String()).Msg("play step")

		diff := step.Offset.Sub(offset)
		if diff.Positive() {
			xRange := function.NewRange(offset, step.Offset)

			timeDelta, err := fs.TimeDelta(xRange, settings.TempoSamplePeriod)
			if err != nil {
				return fmt.Errorf("failed to compute time delta: %w", err)
			}

			log.Debug().Dur("dur", timeDelta).Msg("sleeping")
			time.Sleep(timeDelta)
		}

		for i, events := range step.TrackEvents {
			t := tracks[i]

			wr.SetChannel(t.Channel)

			for _, e := range events {
				e.Write(wr)
			}
		}

		offset = step.Offset
	}

	log.Info().Msg("MIDI playback done")

	return nil
}

func makePlaySteps(tracks []*Track) []*playStep {
	offsets := rat.Rats{}

	// gather all the unique offsets
	for _, track := range tracks {
		offsets = offsets.Union(track.Events.Offsets())
	}

	sort.Sort(offsets)

	steps := make([]*playStep, len(offsets))

	for i, offset := range offsets {
		trackEvents := map[int]NoteEvents{}

		for j, track := range tracks {
			events := track.Events.WithOffset(offset)

			if len(events) > 0 {
				trackEvents[j] = events
			}
		}

		steps[i] = &playStep{
			Offset:      offset,
			TrackEvents: trackEvents,
		}
	}

	return steps
}
