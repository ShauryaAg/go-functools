package tools

// Find returns first element index in 'slice'
// that satisfies the given 'callback' function.
func FindIndex[T any](slice []T, callback func(T) bool) int {
	for i, val := range slice {
		if callback(val) {
			return i
		}
	}
	return -1
}
