package tests

import (
	"testing"
	tools "tools/tools"
)

type filterTest[T any] struct {
	slice []T
	callback func(T) bool
	expected []T
}

var intFilterTest = []filterTest[int] {
	filterTest[int] {[]int{1, 2, 3, 4, 5}, func(val int) bool { return val > 2 }, []int{3, 4, 5}},
	filterTest[int] {[]int{1, 2, 3, 4, 5}, func(val int) bool { return val < 2 }, []int{1}},
	filterTest[int] {[]int{1, 2, 3, 4, 5}, func(val int) bool { return val == 2 }, []int{2}},
}

var stringFilterTest = []filterTest[string] {
	filterTest[string] {[]string{"a", "b", "c", "d", "e"}, func(val string) bool { return val == "c" }, []string{"c"}},
	filterTest[string] {[]string{"a", "b", "c", "d", "e"}, func(val string) bool { return val == "f" }, []string{}},
}

func TestFilter(t *testing.T) {
	for idx, test := range intFilterTest {
		actual := tools.Filter(test.slice, test.callback)
		if tools.ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. Filter() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}

	for idx, test := range stringFilterTest {
		actual := tools.Filter(test.slice, test.callback)
		if tools.ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. Filter() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}
}
