//go:build goexperiment.rangefunc

package rangeref_test

import (
	"fmt"

	"github.com/rodaine/rangeref"
)

func Example() {
	x := make([]int, 3)
	for i, n := range rangeref.Slice(x) {
		*n = i * 2
	}
	fmt.Println(x)
	// Output:
	// [0 2 4]
}
