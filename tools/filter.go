package tools

// Filter finds all the elements in the 
// 'slice' that satisfy the 'callback' function.
func Filter[T any](slice []T, callback func(T) bool) []T {
	var result []T
	for _, val := range slice {
		if callback(val) {
			result = append(result, val)
		}
	}
	return result
}
