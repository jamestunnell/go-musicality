package midismf_test

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/commands/midismf"
	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/jamestunnell/go-musicality/notation/section"
)

const (
	testSectionName = "testSection"
)

func TestLoadScoresEmpty(t *testing.T) {
	scores, err := midismf.LoadScores()

	assert.NoError(t, err)
	assert.Empty(t, scores)
}

func TestLoadScoresNonExistant(t *testing.T) {
	scores, err := midismf.LoadScores(bogusPath(t))

	assert.Error(t, err)
	assert.Empty(t, scores)
}

func TestLoadScoresNotScoreJSON(t *testing.T) {
	notJSON := createTemp(t)

	defer os.Remove(notJSON.Name())

	err := ioutil.WriteFile(notJSON.Name(), []byte("not-json-data"), fs.ModeExclusive)

	require.NoError(t, err)

	scores, err := midismf.LoadScores(notJSON.Name())

	assert.Error(t, err)
	assert.Empty(t, scores)
}

func TestLoadScoresValidOnly(t *testing.T) {
	testLoadScores(t, func(valid, invalid string) {
		scores, err := midismf.LoadScores(valid)
		assert.NoError(t, err)
		assert.NotEmpty(t, scores)
	})
}

func TestLoadScoresInvalidOnly(t *testing.T) {
	testLoadScores(t, func(valid, invalid string) {
		scores, err := midismf.LoadScores(invalid)
		assert.NoError(t, err)
		assert.NotEmpty(t, scores)
	})
}

func TestLoadScoresValidAndInvalid(t *testing.T) {
	testLoadScores(t, func(valid, invalid string) {
		scores, err := midismf.LoadScores(valid, invalid)
		assert.NoError(t, err)
		assert.NotEmpty(t, scores)
	})
}

func testLoadScores(t *testing.T, f func(valid, invalid string)) {
	invalid := createTemp(t)

	defer os.Remove(invalid.Name())

	valid := createTemp(t)

	defer os.Remove(valid.Name())

	err := ioutil.WriteFile(valid.Name(), validScoreJSON(t), fs.ModeExclusive)

	require.NoError(t, err)

	err = ioutil.WriteFile(invalid.Name(), invalidScoreJSON(t), fs.ModeExclusive)

	require.NoError(t, err)

	f(valid.Name(), invalid.Name())
}

func invalidScoreJSON(t *testing.T) []byte {
	s := validScore()

	// invalidate
	s.Sections[testSectionName].Measures[0].Meter.Numerator = 0

	return scoreJSON(t, s)
}

func validScoreJSON(t *testing.T) []byte {
	return scoreJSON(t, validScore())
}

func validScore() *score.Score {
	s := score.New()
	sec := section.New()
	m := measure.New(meter.New(4, 4))

	sec.Measures = append(sec.Measures, m)

	s.Sections[testSectionName] = sec

	return s
}

func scoreJSON(t *testing.T, s *score.Score) []byte {
	d, err := json.Marshal(s)

	require.NoError(t, err)

	return d
}

func createTemp(t *testing.T) *os.File {
	f, err := os.CreateTemp("", "midismf_test*")

	require.NoError(t, err)

	return f
}
