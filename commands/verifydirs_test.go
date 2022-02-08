package commands_test

import (
	"os"
	"testing"

	"github.com/jamestunnell/go-musicality/commands"
	"github.com/jamestunnell/go-musicality/commands/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerifyDirsEmpty(t *testing.T) {
	assert.NoError(t, commands.VerifyDirs())
}

func TestVerifyFilesEmpty(t *testing.T) {
	assert.NoError(t, commands.VerifyFiles())
}

func TestVerify(t *testing.T) {
	dir, err := os.MkdirTemp("", "verifydirs*")

	require.NoError(t, err)

	defer os.Remove(dir)

	f := testutil.CreateTemp(t)

	defer os.Remove(f.Name())

	bogus := testutil.BogusPath(t)

	assert.Error(t, commands.VerifyDirs(bogus))
	assert.Error(t, commands.VerifyDirs(f.Name()))
	assert.NoError(t, commands.VerifyDirs(dir))
	assert.Error(t, commands.VerifyDirs(dir, f.Name()))
	assert.Error(t, commands.VerifyDirs(dir, bogus))

	assert.Error(t, commands.VerifyFiles(bogus))
	assert.Error(t, commands.VerifyFiles(dir))
	assert.NoError(t, commands.VerifyFiles(f.Name()))
	assert.Error(t, commands.VerifyFiles(f.Name(), dir))
	assert.Error(t, commands.VerifyFiles(f.Name(), bogus))
}
