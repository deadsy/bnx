//-----------------------------------------------------------------------------
/*

 */
//-----------------------------------------------------------------------------

package bn

import (
	"fmt"
	"strings"
)

//-----------------------------------------------------------------------------

type Pmf map[string]float32

func NewPmf() Pmf {
	return Pmf{}
}

func (pmf Pmf) Set(label string, x float32) {
	pmf[label] = x
}

func (pmf Pmf) String() string {
	s := make([]string, len(pmf))
	i := 0
	for k, v := range pmf {
		s[i] = fmt.Sprintf("%s: %f", k, v)
		i += 1
	}
	return strings.Join(s, "\n")
}

//-----------------------------------------------------------------------------
