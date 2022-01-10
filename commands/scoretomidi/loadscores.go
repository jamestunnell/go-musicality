package scoretomidi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/rs/zerolog/log"
)

type Scores map[string]*score.Score

func LoadScores(scorePaths ...string) (map[string]*score.Score, error) {
	scores := Scores{}

	for _, fpath := range scorePaths {
		f, err := os.Open(fpath)
		if err != nil {
			err = fmt.Errorf("failed to open score file '%s': %w", fpath, err)

			return Scores{}, err
		}

		d, err := ioutil.ReadAll(f)
		if err != nil {
			err = fmt.Errorf("failed to read score file '%s': %w", fpath, err)

			return Scores{}, err
		}

		var s score.Score

		if err = json.Unmarshal(d, &s); err != nil {
			err = fmt.Errorf("failed to unmarshal score file '%s': %w", fpath, err)

			return Scores{}, err
		}

		log.Info().Str("path", fpath).Msg("loaded score")

		scores[fpath] = &s
	}

	return scores, nil
}
