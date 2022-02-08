package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/rs/zerolog/log"
	"github.com/xeipuuv/gojsonschema"
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

		// before unmarshaling, verify with JSON schema
		result, err := score.Schema().Validate(gojsonschema.NewBytesLoader(d))
		if err != nil {
			err = fmt.Errorf("failed to validate score file '%s': %w", fpath, err)

			return Scores{}, err
		}

		if !result.Valid() {
			details := &strings.Builder{}
			for _, err := range result.Errors() {
				details.WriteString("\n - ")
				details.WriteString(err.String())
			}

			err := fmt.Errorf("score file '%s' is not valid: %s", fpath, details.String())

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
