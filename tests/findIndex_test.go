package tests

import (
	"testing"

	tools "github.com/ShauryaAg/go-functools/tools"
)

type findIndexTest[T any] struct {
	slice []T
	callback func(T) bool
	expected int
}

var intFindIndexTest = []findIndexTest[int] {
	findIndexTest[int] {[]int{1, 2, 3, 4, 5}, func(val int) bool { return val > 2 }, 2},
	findIndexTest[int] {[]int{1, 2, 3, 4, 5}, func(val int) bool { return val < 2 }, 0},
	findIndexTest[int] {[]int{1, 2, 3, 4, 5}, func(val int) bool { return val == 2 }, 1},
}

var stringFindIndexTest = []findIndexTest[string] {
	findIndexTest[string] {[]string{"a", "b", "c", "d", "e"}, func(val string) bool { return val == "c" }, 2},
	findIndexTest[string] {[]string{"a", "b", "c", "d", "e"}, func(val string) bool { return val == "f" }, -1},
}

func TestFindIndex(t *testing.T) {
	for idx, test := range intFindIndexTest {
		actual := tools.FindIndex(test.slice, test.callback)
		if actual != test.expected {
			t.Errorf("%d. FindIndex() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}

	for idx, test := range stringFindIndexTest {
		actual := tools.FindIndex(test.slice, test.callback)
		if actual != test.expected {
			t.Errorf("%d. FindIndex() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}
}