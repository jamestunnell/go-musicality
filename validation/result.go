package validation

type Result struct {
	Context    string
	Errors     []error
	SubResults []*Result
}
