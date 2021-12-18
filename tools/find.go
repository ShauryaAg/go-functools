package tools 

// Find the first element in an array that satisfies the given condition.
func Find[T any] (slice []T, callback func(T) bool) T {	
	for _, val := range slice {
		if callback(val) {
			return val
		}
	}
	return nil
}