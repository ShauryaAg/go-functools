package tools

func reverse[T any] (slice []T, callback func(T) bool) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
        slice[i], slice[j] = slice[j], slice[i]
    }

	return slice
}