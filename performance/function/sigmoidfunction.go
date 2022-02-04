package function

import (
	"math"
)

var (
	SigmoidDomain = NewRange(-5, 5)
	SigmoidRange  = NewRange(sigmoid(SigmoidDomain.Start), sigmoid(SigmoidDomain.End))
	SigmoidSpan   = SigmoidRange.End - SigmoidRange.Start
)

type SigmoidFunction struct {
	domain Range
	y0     float64
	dy     float64
}

func NewSigmoidFunction(p0, p1 Point) *SigmoidFunction {
	if p0.X > p1.X {
		return NewSigmoidFunction(p1, p0)
	}

	dy := p1.Y - p0.Y
	domain := NewRange(p0.X, p1.X)

	return &SigmoidFunction{domain: domain, y0: p0.Y, dy: dy}
}

func (f *SigmoidFunction) At(x float64) float64 {
	x_ := transformDomains(f.domain, SigmoidDomain, x)
	y_ := (sigmoid(x_) - SigmoidRange.Start) / SigmoidSpan
	return f.y0 + y_*f.dy
}

func (f *SigmoidFunction) Domain() Range {
	return f.domain
}

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func invertedSigmoid(y float64) float64 {
	return -math.Log((1 - y) / y)
}

// transformDomains moves from x in start domain, transformed to x in end domain
func transformDomains(startDomain, endDomain Range, x float64) float64 {
	perc := (x - startDomain.Start) / (startDomain.End - startDomain.Start)
	return perc*(endDomain.End-endDomain.Start) + endDomain.Start
}

// 		#def from(y)
// 		#	y2 = (y - @y0) / @dy
// 		#	x2 = Sigmoid.inv_sigm(y2 * SIGM_SPAN + SIGM_RANGE.Start)
// 		#	x = Function.transform_domains(SIGM_DOMAIN, @external_domain, x2)
// 		#	return x
// 		#end

// 		# Given a domain, an xy-point in that domain, and the y-value at
// 		# the end of the domain, find the y-value at the start of the domain,
// 		# assuming the the function is sigmoid.
// 		def self.find_y0 domain, pt, y1
// 			x,y = pt
// 			x_ = Function.transform_domains(domain, SIGM_DOMAIN, x)
// 			y_ = (sigm(x_) - SIGM_RANGE.Start) / SIGM_SPAN
// 			return Function::Linear.new([y_,y],[1,y1]).at(0)
// 		end
