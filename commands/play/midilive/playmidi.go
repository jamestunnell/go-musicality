package midilive

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/go-musicality/commands"
	"github.com/jamestunnell/go-musicality/performance/midi"
)

type PlayMIDI struct {
	ScoreFile string
}

const Name = "midi-live"

var errNoScoreFiles = errors.New("no score files")

func NewFromArgs(cliArgs ...string) (*PlayMIDI, error) {
	var scoreFile string

	switch len(cliArgs) {
	case 0:
		return nil, fmt.Errorf("no score file argument")
	case 1:
		scoreFile = cliArgs[0]
	default:
		return nil, fmt.Errorf("too many arguments")
	}

	cmd := &PlayMIDI{
		ScoreFile: scoreFile,
	}

	return cmd, nil
}

func (cmd *PlayMIDI) Name() string {
	return Name
}

func (cmd *PlayMIDI) Execute() error {
	if err := commands.VerifyFiles(cmd.ScoreFile); err != nil {
		return fmt.Errorf("failed to verify score file: %w", err)
	}

	scores, err := commands.LoadScores(true, cmd.ScoreFile)
	if err != nil {
		return fmt.Errorf("failed to load score: %w", err)
	}

	// Make sure all the score is valid
	for fpath, score := range scores {
		if result := score.Validate(); result != nil {
			return fmt.Errorf("score is invalid: %w", result)
		}

		log.Info().
			Str("input", fpath).
			Msg("playing score")

		err := midi.PlayMIDI(score)
		if err != nil {
			return fmt.Errorf("failed to play score: %w", err)
		}
	}

	return nil
}
