package main

import (
	"fmt"

	"github.com/deadsy/bnx/bn"
)

func main() {

	pmf := bn.NewPmf()

	pmf.Set("1", 0.1)
	pmf.Set("2", 0.2)
	pmf.Set("3", 0.3)
	pmf.Set("4", 0.4)

	fmt.Printf("%v\n", pmf)
}
