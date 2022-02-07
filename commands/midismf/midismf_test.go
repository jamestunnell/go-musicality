package midismf_test

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/commands/midismf"
)

func TestScoreToMIDINoScores(t *testing.T) {
	cmd := midismf.ScoreToMIDI{
		OutDir:     "",
		ScoreFiles: []string{},
	}

	assert.NotEmpty(t, cmd.Name())
	assert.Error(t, cmd.Execute())
}

func TestScoreToMIDIScoreNotFound(t *testing.T) {
	cmd := midismf.ScoreToMIDI{
		OutDir:     "",
		ScoreFiles: []string{bogusPath(t)},
	}

	assert.NotEmpty(t, cmd.Name())
	assert.Error(t, cmd.Execute())
}

func TestScoreToMIDIOutDirNotFound(t *testing.T) {
	testLoadScores(t, func(valid, invalid string) {
		cmd := midismf.ScoreToMIDI{
			OutDir:     bogusPath(t),
			ScoreFiles: []string{valid},
		}

		assert.NotEmpty(t, cmd.Name())
		assert.Error(t, cmd.Execute())
	})
}

func TestScoreToMIDIInvalidScore(t *testing.T) {
	testLoadScores(t, func(valid, invalid string) {
		cmd := midismf.ScoreToMIDI{
			OutDir:     "",
			ScoreFiles: []string{invalid},
		}

		assert.Error(t, cmd.Execute())
	})
}

func TestScoreToMIDINotScoreJSON(t *testing.T) {
	notJSON := createTemp(t)

	defer os.Remove(notJSON.Name())

	err := ioutil.WriteFile(notJSON.Name(), []byte("not-json-data"), fs.ModeExclusive)

	require.NoError(t, err)

	cmd := midismf.ScoreToMIDI{
		OutDir:     "",
		ScoreFiles: []string{notJSON.Name()},
	}

	assert.Error(t, cmd.Execute())
}

func TestScoreToMIDIValidAndInvalidScore(t *testing.T) {
	testLoadScores(t, func(valid, invalid string) {
		cmd := midismf.ScoreToMIDI{
			OutDir:     "",
			ScoreFiles: []string{valid, invalid},
		}

		assert.Error(t, cmd.Execute())
	})
}

func TestScoreToMIDIValidScore(t *testing.T) {
	testLoadScores(t, func(valid, invalid string) {
		cmd := midismf.ScoreToMIDI{
			OutDir:     "",
			ScoreFiles: []string{valid},
		}

		assert.NoError(t, cmd.Execute())
	})
}

func bogusPath(t *testing.T) string {
	wd, err := os.Getwd()

	require.NoError(t, err)

	return filepath.Join(wd, "bogus")
}
