package tools

// Run a function for each element in the array, and reduce it into
// a single value, using the given callback function.
func Reduce[T any] (slice []T, callback func(T, T) T, initial T) T {
	var result T = initial
	for _, val := range slice {
		result = callback(result, val)
	}
	return result
}