//-----------------------------------------------------------------------------
/*

Probability Mass Function

*/
//-----------------------------------------------------------------------------

package bn

import (
	"fmt"
	"strings"
)

//-----------------------------------------------------------------------------

// Pmf is a probability mass function.
type Pmf map[string]float32

// NewPmf returns a PMF.
func NewPmf() Pmf {
	return Pmf{}
}

// Set sets a label to a probability value.
func (pmf Pmf) Set(k string, v float32) {
	pmf[k] = v
}

// Mul multiples the value of a label.
func (pmf Pmf) Mul(k string, v float32) {
	pmf[k] *= v
}

// Inc increments the value associated with a PMF label.
func (pmf Pmf) Inc(k string) {
	pmf[k] += 1.0
}

// Normalize scales PMF values so they add to 1.
func (pmf Pmf) Normalize() {
	sum := float32(0.0)
	for _, v := range pmf {
		if v <= 0 {
			panic("bad pmf value")
		}
		sum += v
	}
	if sum <= 0 {
		panic("bad pmf sum")
	}
	for k := range pmf {
		pmf[k] /= sum
	}
}

// P returns the probability of a label.
func (pmf Pmf) P(k string) float32 {
	return pmf[k]
}

// String returns a string for the PMF.
func (pmf Pmf) String() string {
	s := make([]string, len(pmf))
	i := 0
	for k, v := range pmf {
		s[i] = fmt.Sprintf("%s: %f", k, v)
		i++
	}
	return strings.Join(s, "\n")
}

//-----------------------------------------------------------------------------
