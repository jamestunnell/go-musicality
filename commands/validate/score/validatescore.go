package score

import (
	"errors"
	"fmt"

	"github.com/jamestunnell/go-musicality/commands"
)

type ValidateScore struct {
	ScoreFiles []string
}

const Name = "score"

var errNoScoreFiles = errors.New("no score files")

func NewFromArgs(cliArgs ...string) (*ValidateScore, error) {
	cmd := &ValidateScore{
		ScoreFiles: cliArgs,
	}

	return cmd, nil
}

func (cmd *ValidateScore) Name() string {
	return Name
}

func (cmd *ValidateScore) Execute() error {
	if len(cmd.ScoreFiles) == 0 {
		return errNoScoreFiles
	}

	if err := commands.VerifyFiles(cmd.ScoreFiles...); err != nil {
		return fmt.Errorf("failed to verify score files: %w", err)
	}

	for _, fpath := range cmd.ScoreFiles {
		scores, err := commands.LoadScores(false, fpath)
		if err != nil {
			return fmt.Errorf("failed to load '%s'", fpath)
		}

		result := scores[fpath].Validate()
		if result != nil {
			fmt.Printf("'%s' is invalid:\n", fpath)

			for context, err := range result.ContextErrors() {
				fmt.Printf("%s: %v\n", context, err)
			}

			continue
		}

		fmt.Printf("'%s' is valid\n", fpath)
	}

	return nil
}
