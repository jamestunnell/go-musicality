package midismf_test

import (
	"io/fs"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/commands/render/midismf"
	"github.com/jamestunnell/go-musicality/commands/testutil"
)

func TestRenderSMFNoScores(t *testing.T) {
	cmd := midismf.RenderSMF{
		OutDir:     "",
		ScoreFiles: []string{},
	}

	assert.NotEmpty(t, cmd.Name())
	assert.Error(t, cmd.Execute())
}

func TestRenderSMFScoreNotFound(t *testing.T) {
	cmd := midismf.RenderSMF{
		OutDir:     "",
		ScoreFiles: []string{testutil.BogusPath(t)},
	}

	assert.NotEmpty(t, cmd.Name())
	assert.Error(t, cmd.Execute())
}

func TestRenderSMFOutDirNotFound(t *testing.T) {
	scoreJSONs := [][]byte{testutil.ValidScoreJSON(t)}

	testutil.WriteScoreFiles(t, scoreJSONs, func(names []string) {
		cmd := midismf.RenderSMF{
			OutDir:     testutil.BogusPath(t),
			ScoreFiles: names,
		}

		assert.NotEmpty(t, cmd.Name())
		assert.Error(t, cmd.Execute())
	})
}

func TestRenderSMFInvalidScore(t *testing.T) {
	scoreJSONs := [][]byte{testutil.InvalidScoreJSON(t)}

	testutil.WriteScoreFiles(t, scoreJSONs, func(names []string) {
		cmd := midismf.RenderSMF{
			OutDir:     "",
			ScoreFiles: names,
		}

		assert.Error(t, cmd.Execute())
	})
}

func TestRenderSMFNotScoreJSON(t *testing.T) {
	notJSON := testutil.CreateTemp(t)

	defer os.Remove(notJSON.Name())

	err := ioutil.WriteFile(notJSON.Name(), []byte("not-json-data"), fs.ModeExclusive)

	require.NoError(t, err)

	cmd := midismf.RenderSMF{
		OutDir:     "",
		ScoreFiles: []string{notJSON.Name()},
	}

	assert.Error(t, cmd.Execute())
}

func TestRenderSMFValidAndInvalidScore(t *testing.T) {
	scoreJSONs := [][]byte{
		testutil.ValidScoreJSON(t),
		testutil.InvalidScoreJSON(t),
	}

	testutil.WriteScoreFiles(t, scoreJSONs, func(names []string) {
		cmd := midismf.RenderSMF{
			OutDir:     "",
			ScoreFiles: names,
		}

		assert.Error(t, cmd.Execute())
	})
}

func TestRenderSMFHappyPath(t *testing.T) {
	scoreJSONs := [][]byte{testutil.ValidScoreJSON(t)}

	testutil.WriteScoreFiles(t, scoreJSONs, func(names []string) {
		cmd := midismf.RenderSMF{
			OutDir:     "",
			ScoreFiles: names,
		}

		assert.NoError(t, cmd.Execute())
	})
}
