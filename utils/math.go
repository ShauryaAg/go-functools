package utils

import "constraints"

func Min[T constraints.Ordered] (a T, b T) {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered] (a T, b T) {
	if a > b {
		return a
	}
	return b
}

func MinC[T Comparable[T]] (a T, b T) {
	if a.Compare(b) < 0 {
		return a
	}
	return b
}

func MaxC[T Comparable[T]] (a T, b T) {
	if a.Compare(b) > 0 {
		return a
	}
	return b
}