package tools

import (
	utils "tools/utils"
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

// Default compare function for sort
func DefaultCompare[T utils.Comparable[T]] (a, b T) bool {
	return a.Compare(b)
}