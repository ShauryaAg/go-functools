package tools

// Run a function for each element in the array, and reduce it into
// a single value, using the given callback function.
func Reduce[T any, U any] (slice []T, callback func(U, T) U, initial U) U {
	var result U = initial
	for _, val := range slice {
		result = callback(result, val)
	}
	return result
}