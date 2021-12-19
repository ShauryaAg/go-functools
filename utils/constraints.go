package utils

import "constraints"

// Number or String
type NumberOrString interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~complex64 | ~complex128 | ~string
}


// Comparable type
// Should either be Ordered or 
// a type that implements the Compare() method
type Comparable[T any] interface {
	constraints.Ordered
	Compare(T) bool // return 1 if a > b, 0 if a == b, -1 if a < b
}
