package score

import (
	"fmt"

	"gitlab.com/gomidi/midi/smf/smfwriter"
	"gitlab.com/gomidi/midi/writer"
)

func (s *Score) WriteSMF(fpath string) error {
	numTracks := uint16(1)
	write := func(wr *writer.SMF) error {
		return nil
	}

	opts := []smfwriter.Option{smfwriter.NoRunningStatus()}

	err := writer.WriteSMF(fpath, numTracks, write, opts...)
	if err != nil {
		return fmt.Errorf("failed to write SMF: %w", err)
	}

	return nil
}
