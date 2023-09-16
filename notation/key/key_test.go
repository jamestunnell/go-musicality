package key_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/key"
)

func TestMajorKey(t *testing.T) {
	testValidKey(t, key.NewMajor("C"))
	testValidKey(t, key.NewMajor("Bb"))
	testValidKey(t, key.NewMajor("F#"))

	testInvalidKey(t, key.NewMajor(""))
	testInvalidKey(t, key.NewMajor("bogus"))
}

func TestMinorKey(t *testing.T) {
	testValidKey(t, key.NewMinor("C"))
	testValidKey(t, key.NewMinor("Bb"))
	testValidKey(t, key.NewMinor("F#"))

	testInvalidKey(t, key.NewMinor(""))
	testInvalidKey(t, key.NewMinor("bogus"))
}

func testValidKey(t *testing.T, k *key.Key) {
	assert.Nil(t, k.Validate())
}

func testInvalidKey(t *testing.T, k *key.Key) {
	result := k.Validate()

	require.NotNil(t, result)

	assert.NotEmpty(t, result.Errors)
}
