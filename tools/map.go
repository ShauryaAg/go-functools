package tools

// Apply a function to all the elements in an array
func Map[T any, U any](slice []T, callback func(T) U) []U {
	var result []U
	for _, val := range slice {
		result = append(result, callback(val))
	}
	return result
}
