package midismf

import (
	"errors"
	"fmt"
	"os"
)

// VerifyDirs checks that the given dirs exists and are not files.
func VerifyDirs(dirs ...string) error {
	for _, dir := range dirs {
		info, err := os.Stat(dir)
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("dir '%s' does not exist", dir)
		}

		if !info.IsDir() {
			return fmt.Errorf("dir '%s' is a file", dir)
		}

		// log.Info().Str("path", cmd.OutDir).Msg("directory found")
	}

	return nil
}
