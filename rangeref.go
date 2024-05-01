//go:build goexperiment.rangefunc

package rangeref

import (
	"iter"
)

// Slice returns an iter.Seq that ranges over pointers to the values of the
// provided slice. This allows iterating over a slice of values without copying
// or using a slice of pointers instead. This is equivalent to:
//
//	for i := range s {
//	  v := &s[i]
//	  // ...
//	}
func Slice[T any](s []T) iter.Seq2[int, *T] {
	return func(yield func(int, *T) bool) {
		for i := range s {
			if !yield(i, &s[i]) {
				return
			}
		}
	}
}
