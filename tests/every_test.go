package tests

import (
	"testing"

	tools "tools/tools"
)

type everyTest[T any] struct {
	slice []T
	callback func(T) bool
	expected bool
}

var intEveryTests = []everyTest[int] {
	everyTest[int] {[]int{1, 2, 3, 4, 5}, func(i int) bool { return i < 6 }, true},
	everyTest[int] {[]int{1, 2, 3, 4, 5}, func(i int) bool { return i > 3 }, false},
}

var stringEveryTests = []everyTest[string] {
	everyTest[string] {[]string{"a", "b", "c", "d", "e"}, func(i string) bool { return i < "z" }, true},
	everyTest[string] {[]string{"a", "b", "c", "d", "e"}, func(i string) bool { return i == "c" }, false},
}

func TestEvery(t *testing.T) {
	for idx, test := range intEveryTests {
		actual := tools.Every(test.slice, test.callback)
		if actual != test.expected {
			t.Errorf("%d. Every() returned %v, expected %v", idx, actual, test.expected)
		}
	}

	for idx, test := range stringEveryTests {
		actual := tools.Every(test.slice, test.callback)
		if actual != test.expected {
			t.Errorf("%d. Every() returned %v, expected %v", idx, actual, test.expected)
		}
	}
}