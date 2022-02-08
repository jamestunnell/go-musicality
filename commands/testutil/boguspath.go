package testutil

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func BogusPath(t *testing.T) string {
	wd, err := os.Getwd()

	require.NoError(t, err)

	return filepath.Join(wd, "bogus")
}
