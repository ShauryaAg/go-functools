package tests

import (
	"testing"

	tools "github.com/ShauryaAg/go-functools/tools"
)

type reduceTest[T any, U any] struct {
	slice []T
	callback func(U, T) U
	initial U
	expected U
}

var intReduceTest = []reduceTest[int, int] {
	reduceTest[int, int] {[]int{1, 2, 3, 4, 5}, func(val1 int, val2 int) int { return val1 + val2 }, 0, 15},
	reduceTest[int, int] {[]int{1, 2, 3, 4, 5}, func(val1 int, val2 int) int { return val1 * val2 }, 1, 120},
	reduceTest[int, int] {[]int{1, 2, 3, 4, 5}, func(val1 int, val2 int) int { return val1 * val2 }, 2, 240},
}

var stringReduceTest = []reduceTest[string, string] {
	reduceTest[string, string] {[]string{"a", "b", "c", "d", "e"}, func(val1 string, val2 string) string { return val1 + val2 }, "", "abcde"},
}

var runeIntReduceTest = []reduceTest[rune, int] {
	reduceTest[rune, int] {[]rune{'a', 'b', 'c', 'd', 'e'}, func(val1 int, val2 rune) int { return val1 + int(val2) }, 0, 97+98+99+100+101},
}

func TestReduce(t *testing.T) {
	for idx, test := range intReduceTest {
		actual := tools.Reduce(test.slice, test.callback, test.initial)
		if actual != test.expected {
			t.Errorf("%d. Reduce() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}

	for idx, test := range stringReduceTest {
		actual := tools.Reduce(test.slice, test.callback, test.initial)
		if actual != test.expected {
			t.Errorf("%d. Reduce() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}

	for idx, test := range runeIntReduceTest {
		actual := tools.Reduce(test.slice, test.callback, test.initial)
		if actual != test.expected {
			t.Errorf("%d. Reduce() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}
}