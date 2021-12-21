package tools

// Every returns true if all elements in the 'slice' 
// satisfies the test implemented by the 'callback' function.
func Every[T any](slice []T, callback func(T) bool) bool {
	for _, val := range slice {
		result := callback(val)

		if !result {
			return false
		}
	}
	return true
}
