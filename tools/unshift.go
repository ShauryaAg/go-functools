package tools

// Unshift adds 'elems' to the beginning of 
// the 'slice' and returns the new 'slice'.
func Unshift[T any](slice []T, elems ...T) []T {
	return append(elems, slice...)
}
