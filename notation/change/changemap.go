package change

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/jamestunnell/go-musicality/validation"
)

type Map map[*big.Rat]*Change

func (m Map) Validate() *validation.Result {
	results := []*validation.Result{}
	errs := []error{}

	for offset, change := range m {
		if result := change.Validate(); result != nil {
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

func (m Map) SortedOffsets() Rats {
	offsets := make(Rats, len(m))
	i := 0

	for offset := range m {
		offsets[i] = offset

		i++
	}

	sort.Sort(offsets)

	return offsets
}
