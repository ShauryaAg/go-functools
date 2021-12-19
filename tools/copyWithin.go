package tools

import (
	utils "github.com/ShauryaAg/go-functools/utils"
)

// Replace slice[target:target+length] with slice[start:end]
func CopyWithin[T any] (slice []T, target int, start int, end int) []T {
	len := len(slice)

	to := utils.Min(target, len)
	if target < 0 {
		to = utils.Max(len+target, 0)
	}
	from := utils.Min(start, len)
	if start < 0 {
		from = utils.Max(len+start, 0)
	}
	final := utils.Min(end, len)
	if end < 0 {
		final = utils.Max(len+end, 0)
	}
	count := utils.Min(final-from, len-to)

	direction := 1
	if from < to && to < from+count {
		direction = -1
		from += count - 1
		to += count - 1
	}

	// Equivalent to while (count > 0) loop
	for ok := true; ok; ok = (count > 0) {
		slice[int(to)] = slice[int(from)]

		from += direction
		to += direction
		count--
	}

	return slice
}
