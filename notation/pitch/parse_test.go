package pitch_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/pitch"
)

// func TestParseSemitone(t *testing.T) {
// 	// s, err := pitch.ParseSemitone("C#")

// 	// require.NoError(t, err)
// 	// assert.Equal(t, 1, s)

// 	// s, err = pitch.ParseSemitone("C")

// 	// require.NoError(t, err)
// 	// assert.Equal(t, 0, s)

// 	s, err := pitch.ParseSemitone("D b")

// 	require.NoError(t, err)
// 	assert.Equal(t, 1, s)
// }

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
