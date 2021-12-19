package tests

import (
	"testing"

	tools "tools/tools"
)

type shiftTest[T any] struct {
	slice []T
	expected []T
}

var intShiftTests = []shiftTest[int] {
	shiftTest[int] {[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5}},
	shiftTest[int] {[]int{1, 2, 3, 4, 5, 6}, []int{2, 3, 4, 5, 6}},
}

var stringShiftTests = []shiftTest[string] {
	shiftTest[string] {[]string{"a", "b", "c", "d", "e"}, []string{"b", "c", "d", "e"}},
	shiftTest[string] {[]string{"a", "b", "c", "d", "e", "f"}, []string{"b", "c", "d", "e", "f"}},
}

func TestShift(t *testing.T) {
	for idx, test := range intShiftTests {
		actual := tools.Shift(test.slice)
		if tools.ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. Shift() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}

	for idx, test := range stringShiftTests {
		actual := tools.Shift(test.slice)
		if tools.ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. Shift() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}
}