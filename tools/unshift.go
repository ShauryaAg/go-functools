package tools

// Add elements to the beginning of an array
func unshift[T any] (slice []T, elems ...T) []T {
	return append(elems, slice...)
}