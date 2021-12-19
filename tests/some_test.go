package tests

import (
	"testing"

	tools "github.com/ShauryaAg/go-functools/tools"
)

type someTest[T any] struct {
	slice []T
	callback func(T) bool
	expected bool
}

var intSomeTests = []someTest[int] {
	someTest[int] {[]int{1, 2, 3, 4, 5}, func(i int) bool { return i > 3 }, true},
	someTest[int] {[]int{1, 2, 3, 4, 5}, func(i int) bool { return i > 6 }, false},
}

var stringSomeTests = []someTest[string] {
	someTest[string] {[]string{"a", "b", "c", "d", "e"}, func(i string) bool { return i == "c" }, true},
	someTest[string] {[]string{"a", "b", "c", "d", "e"}, func(i string) bool { return i == "z" }, false},
}

func TestSome(t *testing.T) {
	for idx, test := range intSomeTests {
		actual := tools.Some(test.slice, test.callback)
		if actual != test.expected {
			t.Errorf("%d. Some() returned %v, expected %v", idx, actual, test.expected)
		}
	}

	for idx, test := range stringSomeTests {
		actual := tools.Some(test.slice, test.callback)
		if actual != test.expected {
			t.Errorf("%d. Some() returned %v, expected %v", idx, actual, test.expected)
		}
	}
}