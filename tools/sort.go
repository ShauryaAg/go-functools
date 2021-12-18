package tools


// Number or String
type NumberOrString interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~complex64 | ~complex128 | ~string
}


// Comparable type
// Should either NumberOrString or 
// a type that implements the Comparable interface
type Comparable[T any] interface {
	NumberOrString
	Compare(T) bool
}


// Sort an array
func Sort[T any] (slice []T, compareFn func(T, T) bool) []T {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if compareFn(slice[i], slice[j]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	return slice
}

// Default compare function for sort
func DefaultCompare[T Comparable[T]] (a, b T) bool {
	return a.Compare(b)
}