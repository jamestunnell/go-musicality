package midilive_test

import (
	"io/fs"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/commands/play/midilive"
	"github.com/jamestunnell/go-musicality/commands/testutil"
)

func TestPlayMIDIEmptyScore(t *testing.T) {
	cmd := midilive.PlayMIDI{
		ScoreFile: "",
	}

	assert.NotEmpty(t, cmd.Name())
	assert.Error(t, cmd.Execute())
}

func TestPlayMIDIMissingScore(t *testing.T) {
	cmd := midilive.PlayMIDI{
		ScoreFile: testutil.BogusPath(t),
	}

	assert.NotEmpty(t, cmd.Name())
	assert.Error(t, cmd.Execute())
}

func TestPlayMIDIInvalidScore(t *testing.T) {
	scoreJSONs := [][]byte{testutil.InvalidScoreJSON(t)}

	testutil.WriteScoreFiles(t, scoreJSONs, func(names []string) {
		cmd := midilive.PlayMIDI{
			ScoreFile: names[0],
		}

		assert.Error(t, cmd.Execute())
	})
}

func TestPlayMIDINotScoreJSON(t *testing.T) {
	notJSON := testutil.CreateTemp(t)

	defer os.Remove(notJSON.Name())

	err := ioutil.WriteFile(notJSON.Name(), []byte("not-json-data"), fs.ModeExclusive)

	require.NoError(t, err)

	cmd := midilive.PlayMIDI{
		ScoreFile: notJSON.Name(),
	}

	assert.Error(t, cmd.Execute())
}

func TestPlayMIDIHappyPath(t *testing.T) {
	scoreJSONs := [][]byte{testutil.ValidScoreJSON(t)}

	testutil.WriteScoreFiles(t, scoreJSONs, func(names []string) {
		cmd := midilive.PlayMIDI{
			ScoreFile: names[0],
		}

		assert.NoError(t, cmd.Execute())
	})
}
