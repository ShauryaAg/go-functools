package tools

// Remove the first element from the array
func shift[T any] (slice []T) []T {
	return slice[1:]
}