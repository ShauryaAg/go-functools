package tools

// Sort an array
func sort[T comparable] (slice []T, compareFn func(T, T) bool) []T {
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
func defaultCompare[T comparable] (a, b T) bool {
	return a < b
}