package tools

// Remove the first element from the array
func Shift[T any] (slice []T) []T {
	return slice[1:]
}