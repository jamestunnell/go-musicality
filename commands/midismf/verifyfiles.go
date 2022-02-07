package midismf

import (
	"fmt"
	"os"
)

func VerifyFiles(fpaths ...string) error {
	// Make sure all of the score files exist and are not dirs
	for _, fpath := range fpaths {
		info, err := os.Stat(fpath)
		if os.IsNotExist(err) {
			return fmt.Errorf("file '%s' does not exist", fpath)
		}

		if info.IsDir() {
			return fmt.Errorf("file '%s' is a directory", fpath)
		}

		// log.Info().Str("path", fpath).Msg("file found")
	}

	return nil
}
