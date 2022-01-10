package scoretomidi_test

import (
	"os"
	"testing"

	"github.com/jamestunnell/go-musicality/commands/scoretomidi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerifyDirsEmpty(t *testing.T) {
	assert.NoError(t, scoretomidi.VerifyDirs())
}

func TestVerifyFilesEmpty(t *testing.T) {
	assert.NoError(t, scoretomidi.VerifyFiles())
}

func TestVerify(t *testing.T) {
	dir, err := os.MkdirTemp("", "verifydirs*")

	require.NoError(t, err)

	defer os.Remove(dir)

	f := createTemp(t)

	defer os.Remove(f.Name())

	bogus := bogusPath(t)

	assert.Error(t, scoretomidi.VerifyDirs(bogus))
	assert.Error(t, scoretomidi.VerifyDirs(f.Name()))
	assert.NoError(t, scoretomidi.VerifyDirs(dir))
	assert.Error(t, scoretomidi.VerifyDirs(dir, f.Name()))
	assert.Error(t, scoretomidi.VerifyDirs(dir, bogus))

	assert.Error(t, scoretomidi.VerifyFiles(bogus))
	assert.Error(t, scoretomidi.VerifyFiles(dir))
	assert.NoError(t, scoretomidi.VerifyFiles(f.Name()))
	assert.Error(t, scoretomidi.VerifyFiles(f.Name(), dir))
	assert.Error(t, scoretomidi.VerifyFiles(f.Name(), bogus))
}
