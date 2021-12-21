package tools

// Find returns first element in 'slice'
// that satisfies the given 'callback' function.
func Find[T any](slice []T, callback func(T) bool) T {
	for _, val := range slice {
		if callback(val) {
			return val
		}
	}

	// TODO: fix: should return nil
	return slice[0]
}
