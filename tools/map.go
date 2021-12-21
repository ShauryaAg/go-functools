package tools

// Map applies 'callback' function 
// to all the elements in 'slice' and returns the result.
func Map[T any, U any](slice []T, callback func(T) U) []U {
	var result []U
	for _, val := range slice {
		result = append(result, callback(val))
	}
	return result
}
