package tests

import (
	"testing"

	tools "tools/tools"
)

type includesTest[T any]  struct {
	slice []T
	value T
	expected bool
}

var intIncludesTest = []includesTest[int] {
	includesTest[int] {[]int{1, 2, 3, 4, 5}, 2, true},
	includesTest[int] {[]int{1, 2, 3, 4, 5}, 1, true},
	includesTest[int] {[]int{1, 2, 3, 4, 5}, 0, false},
}

var stringIncludesTest = []includesTest[string] {
	includesTest[string] {[]string{"a", "b", "c", "d", "e"}, "c", true},
	includesTest[string] {[]string{"a", "b", "c", "d", "e"}, "f", false},
}

func TestIncludes(t *testing.T) {
	for idx, test := range intIncludesTest {
		actual := tools.Includes(test.slice, test.value)
		if actual != test.expected {
			t.Errorf("%d. Includes() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}

	for idx, test := range stringIncludesTest {
		actual := tools.Includes(test.slice, test.value)
		if actual != test.expected {
			t.Errorf("%d. Includes() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}
}