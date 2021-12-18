package tools


// Find all the elements in an array that satisfy the given condition.
func Filter[T any] (slice []T, callback func(T) bool) []T {
	var result []T
	for _, val := range slice {
		if callback(val) {
			result = append(result, val)
		}
	}
	return result
}