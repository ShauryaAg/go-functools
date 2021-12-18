package tools

// Run a function for each element in the array
func forEach[T any] (slice []T, callback func(T)) {
	for _, val := range slice {
		callback(val)
	}
}