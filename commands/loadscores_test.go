package commands_test

import (
	"io/fs"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/commands"
	"github.com/jamestunnell/go-musicality/commands/testutil"
)

const (
	testSectionName = "testSection"
)

func TestLoadScoresEmpty(t *testing.T) {
	scores, err := commands.LoadScores(true)

	assert.NoError(t, err)
	assert.Empty(t, scores)
}

func TestLoadScoresNonExistant(t *testing.T) {
	scores, err := commands.LoadScores(true, testutil.BogusPath(t))

	assert.Error(t, err)
	assert.Empty(t, scores)
}

func TestLoadScoresNotScoreJSON(t *testing.T) {
	notJSON := testutil.CreateTemp(t)

	defer os.Remove(notJSON.Name())

	err := ioutil.WriteFile(notJSON.Name(), []byte("not-json-data"), fs.ModeExclusive)

	require.NoError(t, err)

	scores, err := commands.LoadScores(true, notJSON.Name())

	assert.Error(t, err)
	assert.Empty(t, scores)
}

func TestLoadScoresValidOnly(t *testing.T) {
	scoreJSONs := [][]byte{
		testutil.ValidScoreJSON(t),
	}

	testutil.WriteScoreFiles(t, scoreJSONs, func(names []string) {
		scores, err := commands.LoadScores(true, names...)

		assert.NoError(t, err)
		assert.NotEmpty(t, scores)
	})
}

func TestLoadScoresInvalidOnly(t *testing.T) {
	scoreJSONs := [][]byte{
		testutil.InvalidScoreJSON(t),
	}

	testutil.WriteScoreFiles(t, scoreJSONs, func(names []string) {
		scores, err := commands.LoadScores(true, names...)

		assert.Error(t, err)
		assert.Empty(t, scores)
	})
}

func TestLoadScoresValidAndInvalid(t *testing.T) {
	scoreJSONs := [][]byte{
		testutil.ValidScoreJSON(t),
		testutil.InvalidScoreJSON(t),
	}

	testutil.WriteScoreFiles(t, scoreJSONs, func(names []string) {
		scores, err := commands.LoadScores(true, names...)

		assert.Error(t, err)
		assert.Empty(t, scores)
	})
}
