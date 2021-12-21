package tools

// Reduce runs a function for each element in 'slice', and reduces it 
// into a single value, using the given 'callback' function and 'initial' value.
func Reduce[T any, U any](slice []T, callback func(U, T) U, initial U) U {
	var result U = initial
	for _, val := range slice {
		result = callback(result, val)
	}
	return result
}
