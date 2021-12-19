package tests

import (
	"testing"

	tools "github.com/ShauryaAg/go-functools/tools"
)

type unshiftTest[T any] struct {
	slice []T
	elems []T
	expected []T
}

var intUnshiftTests = []unshiftTest[int] {
	unshiftTest[int] {[]int{1, 2, 3}, []int{4, 5}, []int{4, 5, 1, 2, 3}},
	unshiftTest[int] {[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
	unshiftTest[int] {[]int{}, []int{4, 5}, []int{4, 5}},
}

var stringUnshiftTests = []unshiftTest[string] {
	unshiftTest[string] {[]string{"a", "b", "c"}, []string{"d", "e"}, []string{"d", "e", "a", "b", "c"}},
	unshiftTest[string] {[]string{"a", "b", "c"}, []string{}, []string{"a", "b", "c"}},
	unshiftTest[string] {[]string{}, []string{"d", "e"}, []string{"d", "e"}},
}

func TestUnshift(t *testing.T) {
	for idx, test := range intUnshiftTests {
		actual := tools.Unshift(test.slice, test.elems...)
		if tools.ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. Sort() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}

	for idx, test := range stringUnshiftTests {
		actual := tools.Unshift(test.slice, test.elems...)
		if tools.ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. Sort() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}
}