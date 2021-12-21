package utils

import "constraints"

// Min returns the minimum value of the two values.
//
// This method is used for Ordered type constraint values
func Min[T constraints.Ordered](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum value of the two values.
//
// This method is used for Ordered type constraint values
func Max[T constraints.Ordered](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

// MinC returns the minimum value of the two values.
//
// This method is used for Comparable type constraint values
func MinC[T Comparable[T]](a T, b T) T {
	if a.Compare(b) < 0 {
		return a
	}
	return b
}

// MaxC returns the maximum value of the two values.
//
// This method is used for Comparable type constraint values
func MaxC[T Comparable[T]](a T, b T) T {
	if a.Compare(b) > 0 {
		return a
	}
	return b
}
