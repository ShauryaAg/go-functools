package tools

import (
	"constraints"

	utils "github.com/ShauryaAg/go-functools/utils"
)

// Sort an array
func Sort[T any] (slice []T, compareFn func(T, T) bool) []T {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if compareFn(slice[i], slice[j]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	return slice
}

func DefaultCompare[T constraints.Ordered] (a, b T) bool {
	return a > b
}

// Default compare function for sort
func DefaultCompareC[T utils.Comparable[T]] (a, b T) bool {
	return a.Compare(b) < 0
}