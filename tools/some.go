package tools

// Check if some of the element(s) in the array satisfy the given condition
func Some[T any](slice []T, callback func(T) bool) bool {
	for _, val := range slice {
		if callback(val) {
			return true
		}
	}
	return false
}
