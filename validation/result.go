package validation

type Result struct {
	Context    string    `json:"context"`
	Errors     []error   `json:"errors"`
	SubResults []*Result `json:"subresults"`
}
