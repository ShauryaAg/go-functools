package tools


// Check if every element in the array satisfies the given condition
func every[T any] (slice []T, callback func(T) bool) bool {
	for _, val := range slice {
		result = callback(val)
		
		if !result {
			return false
		}
	}
	return true
}