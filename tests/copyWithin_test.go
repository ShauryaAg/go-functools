package tests

import (
	"testing"

	tools "tools/tools"
)


type copyWithinTest[T any] struct {
	slice []T
	target int
	start int
	end int
	expected []T
}

var intCopyWithinTests = []copyWithinTest[int] {
	copyWithinTest[int] {[]int{1, 2, 3, 4, 5}, 0, 0, 5, []int{1, 2, 3, 4, 5}},
	copyWithinTest[int] {[]int{1, 2, 3, 4, 5}, 0, 1, 4, []int{2, 3, 4, 4, 5}},
	copyWithinTest[int] {[]int{1, 2, 3, 4, 5}, 0, 2, 3, []int{3, 2, 3, 4, 5}},
}

var stringCopyWithinTests = []copyWithinTest[string] {
	copyWithinTest[string] {[]string{"a", "b", "c", "d", "e"}, 0, 0, 5, []string{"a", "b", "c", "d", "e"}},
	copyWithinTest[string] {[]string{"a", "b", "c", "d", "e"}, 0, 1, 4, []string{"b", "c", "d", "d", "e"}},
	copyWithinTest[string] {[]string{"a", "b", "c", "d", "e"}, 0, 2, 3, []string{"c", "b", "c", "d", "e"}},
}

func TestCopyWithin(t *testing.T) {
	for idx, test := range intCopyWithinTests {
		actual := tools.CopyWithin(test.slice, test.target, test.start, test.end)
		if ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. CopyWithin() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}

	for idx, test := range stringCopyWithinTests {
		actual := tools.CopyWithin(test.slice, test.target, test.start, test.end)
		if ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. CopyWithin() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}
}

