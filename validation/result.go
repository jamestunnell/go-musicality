package validation

import (
	"fmt"
	"strings"
)

type Result struct {
	Context    string    `json:"context"`
	Errors     []error   `json:"errors"`
	SubResults []*Result `json:"subResults"`
}

func (r *Result) Error() string {
	builder := strings.Builder{}

	for ctx, errs := range r.ContextErrors() {
		for _, err := range errs {
			builder.WriteString(fmt.Sprintf("%s: %v\n", ctx, err))
		}
	}

	return builder.String()
}

func (r *Result) ContextErrors() map[string][]error {
	ctxErrs := map[string][]error{}

	if len(r.Errors) > 0 {
		ctxErrs[r.Context] = r.Errors
	}

	for _, result := range r.SubResults {
		for ctx, errs := range result.ContextErrors() {
			ctxErrs[r.Context+"."+ctx] = errs
		}
	}

	return ctxErrs
}
