package testutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateTemp(t *testing.T) *os.File {
	f, err := os.CreateTemp("", "command_test*")

	require.NoError(t, err)

	return f
}
