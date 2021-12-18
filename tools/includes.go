package tools

// Check if the slice includes the given value
func Includes[T comparable] (slice []T, value T) bool {
	for _, val := range slice {
		if val == value {
			return true
		}
	}

	return false
}