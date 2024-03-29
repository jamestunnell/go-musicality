package change

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/validation"
)

type Changes []*Change

func (changes Changes) Len() int {
	return len(changes)
}

func (changes Changes) Swap(i, j int) {
	changes[i], changes[j] = changes[j], changes[i]
}

func (changes Changes) Less(i, j int) bool {
	return rat.IsLess(changes[i].Offset, changes[j].Offset)
}

func (changes Changes) Validate(r ValueRange) *validation.Result {
	results := []*validation.Result{}
	errs := []error{}

	for i, change := range changes {
		if result := change.Validate(r); result != nil {
			result.Context = fmt.Sprintf("%s %d", result.Context, i)
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

// func (m Map) SortedOffsets() rat.Rats {
// 	offsets := make(rat.Rats, len(m))
// 	i := 0

// 	for offset := range m {
// 		offsets[i] = offset

// 		i++
// 	}

// 	sort.Sort(offsets)

// 	return offsets
// }
