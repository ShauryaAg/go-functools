package tools

// Replace slice[start:end] with slice[target]
func copyWithin[T any] (slice []T, target int, start int, end int) []T {
	start = start % len(slice)
	end = end % len(slice)

	for i := start; i < end; i++ {
		slice[i] = slice[target]
	}

	return slice
}