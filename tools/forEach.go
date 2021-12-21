package tools

// ForEach runs 'callback' function 
// for each element in the 'slice'.
func ForEach[T any](slice []T, callback func(T)) {
	for _, val := range slice {
		callback(val)
	}
}
