package tools

// Join the given array with the given separator
func join[T any] (slice []T, separator string) string {
	var result string

	for i, val := range slice {
		if i > 0 {
			result += separator
		}
		result += fmt.Sprintf("%v", val)
	}
	return result
}