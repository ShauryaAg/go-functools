package tools


// Apply a function to all the elements in an array
func map[T any] (slice []T, callback func(T) any) []any {
	var result []any
	for _, val := range slice {
		result = append(result, callback(val))
	}
	return result
}