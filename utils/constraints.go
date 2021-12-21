package utils

import "constraints"

// Number or String
type NumberOrString interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~complex64 | ~complex128 | ~string
}

// Comparable type
// a type that implements the Compare() method
type Comparable[T any] interface {
	Compare(T) int // return 1 if a > b, 0 if a == b, -1 if a < b
}

// Comparable or Ordered type
// a type that either implements the Comparable or Ordered interface
type ComparableOrOrdered[T any] interface {
	Comparable[T] | constraints.Ordered
}