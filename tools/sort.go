package tools

import (
	"constraints"

	utils "github.com/ShauryaAg/go-functools/utils"
)

// Sort the 'slice' using the 'compareFn' function.
func Sort[T any](slice []T, compareFn func(T, T) bool) []T {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if compareFn(slice[i], slice[j]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	return slice
}

// Default compare function for Sort[T] with Ordered type constraint.
func DefaultCompare[T constraints.Ordered](a, b T) bool {
	return a > b
}

// Default compare function for Sort[T] with Comparable type constraint.
//
// Comparable type constraint is a custom constraint that checks the type 
// to contain a Compare() method.
func DefaultCompareC[T utils.Comparable[T]](a, b T) bool {
	return a.Compare(b) < 0
}
