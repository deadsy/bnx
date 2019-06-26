package main

import (
	"fmt"

	"github.com/deadsy/bnx/bn"
)

func main() {

	pmf := bn.NewPmf()
	pmf.Set("Bowl 1", 0.5)
	pmf.Set("Bowl 2", 0.5)

	pmf.Mul("Bowl 1", 0.75)
	pmf.Mul("Bowl 2", 0.5)

	pmf.Normalize()

	fmt.Printf("%f\n", pmf.P("Bowl 1"))
}
