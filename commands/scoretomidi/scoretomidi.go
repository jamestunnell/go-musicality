package scoretomidi

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/go-musicality/performance/midi"
)

type ScoreToMIDI struct {
	ScoreFiles []string
	OutDir     string
}

var errNoScoreFiles = errors.New("no score files")

func NewFromArgs(cliArgs ...string) (*ScoreToMIDI, error) {
	flagSet := flag.NewFlagSet("midi", flag.ExitOnError)
	outDir := flagSet.String("outdir", "", "output directory")

	if err := flagSet.Parse(os.Args[2:]); err != nil {
		return nil, fmt.Errorf("failed to parse midi args: %w", err)
	}

	cmd := &ScoreToMIDI{
		ScoreFiles: flagSet.Args(),
		OutDir:     *outDir,
	}

	return cmd, nil
}

func (cmd *ScoreToMIDI) Name() string {
	return "score-to-MIDI"
}

func (cmd *ScoreToMIDI) Execute() error {
	if len(cmd.ScoreFiles) == 0 {
		return errNoScoreFiles
	}

	if err := VerifyFiles(cmd.ScoreFiles...); err != nil {
		return fmt.Errorf("failed to verify score files: %w", err)
	}

	if cmd.OutDir != "" {
		if err := VerifyDirs(cmd.OutDir); err != nil {
			return fmt.Errorf("failed to verify output dir: %w", err)
		}
	}

	scores, err := LoadScores(cmd.ScoreFiles...)
	if err != nil {
		return fmt.Errorf("failed to load scores: %w", err)
	}

	// Make sure all the scores are valid
	for fpath, score := range scores {
		if result := score.Validate(); result != nil {
			return fmt.Errorf("score '%s' is invalid: %w", fpath, result)
		}
	}

	// Convert the scores to MIDI
	for fpath, s := range scores {
		ext := filepath.Ext(fpath)
		base := filepath.Base(fpath)
		baseWithoutExt := strings.Replace(base, ext, "", 1)

		outDir := cmd.OutDir
		if cmd.OutDir == "" {
			outDir = filepath.Dir(fpath)
		}

		midiFPath := filepath.Join(outDir, baseWithoutExt+".mid")

		log.Info().
			Str("input", fpath).
			Str("output", midiFPath).
			Msg("converting score")

		if err = midi.WriteSMF(s, midiFPath); err != nil {
			return fmt.Errorf("failed to convert '%s' to MIDI: %w", fpath, err)
		}
	}

	return nil
}
