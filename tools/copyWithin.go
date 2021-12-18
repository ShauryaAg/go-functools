package tools

import (
	math "math"
)

// Replace slice[start:end] with slice[target]
func CopyWithin[T any] (slice []T, target int, start int, end int) []T {
	len := len(slice)

	to := math.Min(float64(target), float64(len))
	if target < 0 {
		to = math.Max(float64(len+target), 0)
	}
	from := math.Min(float64(start), float64(len))
	if start < 0 {
		from = math.Max(float64(len+start), 0)
	}
	final := math.Min(float64(end), float64(len))
	if end < 0 {
		final = math.Max(float64(len+end), 0)
	}
	count := math.Min(final-from, float64(len)-to)

	direction := 1.0
	if from < to && to < from+count {
		direction = -1
		from += count - 1
		to += count - 1
	}

	for ok := true; ok; ok = (count > 0) {
		slice[int(to)] = slice[int(from)]

		from += direction
		to += direction
		count--
	}

	return slice
}
