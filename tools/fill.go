package tools

func Fill[T any](slice []T, value T, start int, end int) []T {
	for i := start; i < end; i++ {
		slice[i] = value
	}

	return slice
}
