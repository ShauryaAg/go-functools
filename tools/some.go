package tools

// Some returns true if some of the element(s) in the 'slice'
// satisfy the given 'callback' function, otherwise false.
func Some[T any](slice []T, callback func(T) bool) bool {
	for _, val := range slice {
		if callback(val) {
			return true
		}
	}
	return false
}
