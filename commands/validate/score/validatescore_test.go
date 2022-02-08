package score_test

import (
	"io/fs"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/commands/testutil"
	validate "github.com/jamestunnell/go-musicality/commands/validate/score"
)

func TestScoreToMIDINoScores(t *testing.T) {
	cmd := &validate.ValidateScore{
		ScoreFiles: []string{},
	}

	assert.NotEmpty(t, cmd.Name())
	assert.Error(t, cmd.Execute())
}

func TestScoreToMIDIScoreNotFound(t *testing.T) {
	cmd := &validate.ValidateScore{
		ScoreFiles: []string{testutil.BogusPath(t)},
	}

	assert.NotEmpty(t, cmd.Name())
	assert.Error(t, cmd.Execute())
}

func TestScoreToMIDINotScoreJSON(t *testing.T) {
	notJSON := testutil.CreateTemp(t)

	defer os.Remove(notJSON.Name())

	err := ioutil.WriteFile(notJSON.Name(), []byte("not-json-data"), fs.ModeExclusive)

	require.NoError(t, err)

	cmd := &validate.ValidateScore{
		ScoreFiles: []string{notJSON.Name()},
	}

	assert.Error(t, cmd.Execute())
}

func TestScoreToMIDIInvalidScore(t *testing.T) {
	scoreJSONs := [][]byte{testutil.InvalidScoreJSON(t)}

	testutil.WriteScoreFiles(t, scoreJSONs, func(names []string) {
		cmd := &validate.ValidateScore{
			ScoreFiles: names,
		}

		assert.NoError(t, cmd.Execute())
	})
}

func TestScoreToMIDIValidAndInvalidScore(t *testing.T) {
	scoreJSONs := [][]byte{
		testutil.ValidScoreJSON(t),
		testutil.InvalidScoreJSON(t),
	}

	testutil.WriteScoreFiles(t, scoreJSONs, func(names []string) {
		cmd := &validate.ValidateScore{
			ScoreFiles: names,
		}

		assert.NoError(t, cmd.Execute())
	})
}

func TestScoreToMIDIValidScore(t *testing.T) {
	scoreJSONs := [][]byte{testutil.ValidScoreJSON(t)}

	testutil.WriteScoreFiles(t, scoreJSONs, func(names []string) {
		cmd := &validate.ValidateScore{
			ScoreFiles: names,
		}

		assert.NoError(t, cmd.Execute())
	})
}
