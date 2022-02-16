package testutil

import (
	"io/fs"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func WriteScoreFiles(t *testing.T, scoreJSONs [][]byte, f func(fnames []string)) {
	files := make([]*os.File, len(scoreJSONs))
	names := make([]string, len(scoreJSONs))

	for i := range scoreJSONs {
		files[i] = CreateTemp(t)
		names[i] = files[i].Name()
	}

	defer func() {
		for _, f := range files {
			os.Remove(f.Name())
		}
	}()

	for i, f := range files {
		err := ioutil.WriteFile(f.Name(), scoreJSONs[i], fs.ModePerm)

		require.NoError(t, err)
	}

	f(names)
}
