package tests

import (
	"testing"

	tools "tools/tools"
)

type reverseTest[T any] struct {
	slice []T
	expected []T
}

var intReverseTests = []reverseTest[int] {
	reverseTest[int] {[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	reverseTest[int] {[]int{1, 2, 3, 4, 5, 6}, []int{6, 5, 4, 3, 2, 1}},
	reverseTest[int] {[]int{1, 2, 3, 4, 5, 6, 7}, []int{7, 6, 5, 4, 3, 2, 1}},
}

var stringReverseTests = []reverseTest[string] {
	reverseTest[string] {[]string{"a", "b", "c", "d", "e"}, []string{"e", "d", "c", "b", "a"}},
	reverseTest[string] {[]string{"a", "b", "c", "d", "e", "f"}, []string{"f", "e", "d", "c", "b", "a"}},
	reverseTest[string] {[]string{"a", "b", "c", "d", "e", "f", "g"}, []string{"g", "f", "e", "d", "c", "b", "a"}},
}

func TestReverse(t * testing.T) {
	for idx, test := range intReverseTests {
		actual := tools.Reverse(test.slice)
		if tools.ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. Reverse() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}

	for idx, test := range stringReverseTests {
		actual := tools.Reverse(test.slice)
		if tools.ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. Reverse() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}
}