package midismf_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/commands/midismf"
)

func TestVerifyDirsEmpty(t *testing.T) {
	assert.NoError(t, midismf.VerifyDirs())
}

func TestVerifyFilesEmpty(t *testing.T) {
	assert.NoError(t, midismf.VerifyFiles())
}

func TestVerify(t *testing.T) {
	dir, err := os.MkdirTemp("", "verifydirs*")

	require.NoError(t, err)

	defer os.Remove(dir)

	f := createTemp(t)

	defer os.Remove(f.Name())

	bogus := bogusPath(t)

	assert.Error(t, midismf.VerifyDirs(bogus))
	assert.Error(t, midismf.VerifyDirs(f.Name()))
	assert.NoError(t, midismf.VerifyDirs(dir))
	assert.Error(t, midismf.VerifyDirs(dir, f.Name()))
	assert.Error(t, midismf.VerifyDirs(dir, bogus))

	assert.Error(t, midismf.VerifyFiles(bogus))
	assert.Error(t, midismf.VerifyFiles(dir))
	assert.NoError(t, midismf.VerifyFiles(f.Name()))
	assert.Error(t, midismf.VerifyFiles(f.Name(), dir))
	assert.Error(t, midismf.VerifyFiles(f.Name(), bogus))
}
