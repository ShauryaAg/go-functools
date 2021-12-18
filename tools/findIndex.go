package tools


// Find the first element index in an array that satisfies the given condition.
func findIndex[T any] (slice []T, callback func(T) bool) int {
	for i, val := range slice {
		if callback(val) {
			return i
		}
	}
	return -1
}