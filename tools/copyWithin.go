package tools

func copyWithin[T any] (slice []T, value T, start int, end int) []T {
	start = start % len(slice)
	end = end % len(slice)

	for i := start; i < end; i++ {
		slice[i] = value
	}

	return slice
}