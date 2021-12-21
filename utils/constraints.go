package utils

// Comparable type
// a type that implements the Compare() method
type Comparable[T any] interface {
	Compare(T) int // return 1 if a > b, 0 if a == b, -1 if a < b
}
