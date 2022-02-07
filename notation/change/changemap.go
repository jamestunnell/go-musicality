package change

import (
	"fmt"
	"sort"

	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/validation"
)

type Map map[rat.Rat]*Change

func (m Map) Validate(r ValueRange) *validation.Result {
	results := []*validation.Result{}
	errs := []error{}

	for offset, change := range m {
		if result := change.Validate(r); result != nil {
			result.Context = fmt.Sprintf("%s at offset %v", result.Context, offset)
			results = append(results, result)
		}
	}

	if len(results) == 0 && len(errs) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "changes",
		Errors:     errs,
		SubResults: results,
	}
}

func (m Map) SortedOffsets() rat.Rats {
	offsets := make(rat.Rats, len(m))
	i := 0

	for offset := range m {
		offsets[i] = offset

		i++
	}

	sort.Sort(offsets)

	return offsets
}
