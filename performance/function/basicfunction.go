package function

// type BasicFunction struct {
// 	at     AtFunc
// 	domain util.Range
// }

// type AtFunc func(float64) float64

// func NewFunction(domain util.Range, f AtFunc) (*Function, error) {
// 	if !domain.IsValid() {
// 		return nil, fmt.Errorf("domain %v is not valid", domain)
// 	}

// 	Function := &Function{
// 		at:     f,
// 		domain: domain,
// 	}

// 	return Function, nil
// }

// func (f *Function) At(x float64) (float64, error) {
// 	if !f.domain.IncludesValue(x) {
// 		return 0.0, fmt.Errorf("domain %v does not include value %f", f.domain, x)
// 	}

// 	return f.at(x), nil
// }
