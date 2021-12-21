package tools

// Shift removes the first element from the 'slice' and returns it.
func Shift[T any](slice []T) []T {
	return slice[1:]
}
