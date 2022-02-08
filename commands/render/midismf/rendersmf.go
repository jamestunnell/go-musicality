package midismf

import (
	"errors"
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/go-musicality/commands"
	"github.com/jamestunnell/go-musicality/performance/midi"
)

type RenderSMF struct {
	ScoreFiles []string
	OutDir     string
}

const Name = "midi-smf"

var errNoScoreFiles = errors.New("no score files")

func NewFromArgs(cliArgs ...string) (*RenderSMF, error) {
	flagSet := flag.NewFlagSet(Name, flag.ExitOnError)
	outDir := flagSet.String("outdir", "", "output directory")

	if err := flagSet.Parse(cliArgs); err != nil {
		return nil, fmt.Errorf("failed to parse midi args: %w", err)
	}

	cmd := &RenderSMF{
		ScoreFiles: flagSet.Args(),
		OutDir:     *outDir,
	}

	return cmd, nil
}

func (cmd *RenderSMF) Name() string {
	return Name
}

func (cmd *RenderSMF) Execute() error {
	if len(cmd.ScoreFiles) == 0 {
		return errNoScoreFiles
	}

	if err := commands.VerifyFiles(cmd.ScoreFiles...); err != nil {
		return fmt.Errorf("failed to verify score files: %w", err)
	}

	if cmd.OutDir != "" {
		if err := commands.VerifyDirs(cmd.OutDir); err != nil {
			return fmt.Errorf("failed to verify output dir: %w", err)
		}
	}

	scores, err := commands.LoadScores(cmd.ScoreFiles...)
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
