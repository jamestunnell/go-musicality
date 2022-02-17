package function

import (
	"errors"
	"fmt"

	"github.com/jamestunnell/go-musicality/common/rat"
)

type SubdomainFunctionPair struct {
	Subdomain Range
	Function  Function
}

type PiecewiseFunction struct {
	domain Range
	pairs  []SubdomainFunctionPair
}

// NewPiecewiseFunction creates a piecewise function using the given
// subdomain-function pairs. The subdomains must be contiguous.
func NewPiecewiseFunction(pairs []SubdomainFunctionPair) (*PiecewiseFunction, error) {
	n := len(pairs)
	if n == 0 {
		return nil, errors.New("no subdomain-function pairs given")
	}

	for i, pair := range pairs {
		d := pair.Subdomain
		f := pair.Function

		if !f.Domain().IncludesRange(d) {
			err := fmt.Errorf("subdomain %v is not included in domain %v", d, f.Domain())
			return nil, err
		}

		// Make sure subdomains are contiguous
		if i > 0 {
			dPrev := pairs[i-1].Subdomain
			if !dPrev.End.Equal(d.Start) {
				err := fmt.Errorf("subdomain %v is not contiguous with prev subdomain %v", d, dPrev)
				return nil, err
			}
		}
	}

	domain := NewRange(pairs[0].Subdomain.Start, pairs[n-1].Subdomain.End)

	return &PiecewiseFunction{domain: domain, pairs: pairs}, nil
}

func (f *PiecewiseFunction) Domain() Range {
	return f.domain
}

func (f *PiecewiseFunction) At(x rat.Rat) float64 {
	var y float64

	n := len(f.pairs)

	for i, pair := range f.pairs {
		subdomain := pair.Subdomain

		// For all but the last subdomain, the subdomains are actually
		// meant to exclude the end so as to not overlap with the following
		// subdomain
		if (i == (n-1) && subdomain.IncludesValue(x)) || subdomain.IncludesValueExcl(x) {
			y = pair.Function.At(x)
			break
		}
	}

	return y
}

// func (f *PiecewiseFunction) AddPiece(newFunction Function) error {
// 	newDomain := newFunction.Domain()

// 	if !newDomain.IsValid() {
// 		return fmt.Errorf("domain %v is not valid", newDomain)
// 	}

// 	// Any existing domains?
// 	if len(f.pieces) == 0 {
// 		f.pieces[newDomain] = newFunction
// 		return nil
// 	}

// 	// Any existing domain include this domain completely?
// 	for existingDomain, existingFunction := range f.pieces {
// 		if existingDomain.IncludesRange(newDomain) {
// 			// Remove the existing domain and insert new domain. If the new domain does not
// 			// completely cover the existing domain, then add pieces of the existing domain
// 			// that are not covered.
// 			delete(f.pieces, existingDomain)

// 			if newDomain.Start > existingDomain.Start {
// 				beforeNewDomain := NewRange(existingDomain.Start, newDomain.Start)
// 				f.pieces[beforeNewDomain] = existingFunction
// 			}

// 			f.pieces[newDomain] = newFunction

// 			if newDomain.End < existingDomain.End {
// 				afterNewDomain := NewRange(newDomain.End, existingDomain.End)
// 				f.pieces[afterNewDomain] = existingFunction
// 			}

// 			return nil
// 		}
// 	}

// 	// Any existing domain(s) included entirely by this new domain?

// 	// Any existing domain included by this domain partially?

// 	return nil
// }
