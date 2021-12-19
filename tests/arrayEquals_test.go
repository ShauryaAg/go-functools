package tests

import (
	"testing"

	tools "tools/tools"
)

type arrayEqualsTest[T any] struct {
	a []T
	b []T
	expected bool
}

var intArrayEqualsTests = []arrayEqualsTest[int] {
	arrayEqualsTest[int] {[]int{1, 2, 3}, []int{1, 2, 3}, true},
	arrayEqualsTest[int] {[]int{1, 2, 3}, []int{1, 2, 4}, false},
	arrayEqualsTest[int] {[]int{1, 2, 3}, []int{}, false},
	arrayEqualsTest[int] {[]int{}, []int{}, true},
}

var stringArrayEqualsTests = []arrayEqualsTest[string] {
	arrayEqualsTest[string] {[]string{"a", "b", "c"}, []string{"a", "b", "c"}, true},
	arrayEqualsTest[string] {[]string{"a", "b", "c"}, []string{"a", "b", "d"}, false},
	arrayEqualsTest[string] {[]string{"a", "b", "c"}, []string{}, false},
	arrayEqualsTest[string] {[]string{}, []string{}, true},
}

func TestArrayEquals(t *testing.T) {
	for idx, test := range intArrayEqualsTests {
		actual := tools.ArrayEquals(test.a, test.b)
		if actual != test.expected {
			t.Errorf("%d. ArrayEquals() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}

	for idx, test := range stringArrayEqualsTests {
		actual := tools.ArrayEquals(test.a, test.b)
		if actual != test.expected {
			t.Errorf("%d. ArrayEquals() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}
}