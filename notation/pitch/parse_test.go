package pitch_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/pitch"
)

func TestParseSemitone(t *testing.T) {
	testParseSemitoneOK(t, "C", 0)
	testParseSemitoneOK(t, "C#", 1)
	testParseSemitoneOK(t, "Db", 1)
	testParseSemitoneFail(t, "C@")
	testParseSemitoneFail(t, "C #")
	testParseSemitoneFail(t, "D b")
	testParseSemitoneFail(t, "H")
	testParseSemitoneFail(t, "")
}

func testParseSemitoneOK(t *testing.T, str string, semitone int) {
	s, err := pitch.ParseSemitone(str)

	require.NoError(t, err)
	assert.Equal(t, semitone, s)
}

func testParseSemitoneFail(t *testing.T, str string) {
	_, err := pitch.ParseSemitone(str)

	require.Error(t, err)
}

func TestParseOK(t *testing.T) {
	testParse(t, "C3", pitch.C3)
	testParse(t, "G5", pitch.G5)
	testParse(t, "Bb13", pitch.New(13, 10))
	testParse(t, "C#4", pitch.Db4)
	testParse(t, "C#2", pitch.Db2)
	testParse(t, "C#0", pitch.Db0)
}

func TestParseFail(t *testing.T) {
	testParseFail(t, " C2")
	testParseFail(t, "C2 ")
	testParseFail(t, "Hb2")
}

func testParseFail(t *testing.T, str string) {
	t.Run(str, func(t *testing.T) {
		_, err := pitch.Parse(str)

		assert.Error(t, err)
	})
}

func testParse(t *testing.T, str string, p *pitch.Pitch) {
	t.Run(str, func(t *testing.T) {
		p2, err := pitch.Parse(str)

		require.NoError(t, err)
		assert.Equal(t, p, p2)
	})
}
