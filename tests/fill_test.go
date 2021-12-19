package tests

import ("testing"
tools "tools/tools")

type fillTest[T any] struct {
	slice []T
	value T
	start int
	end int
	expected []T
}

var intFillTests = []fillTest[int] {
	fillTest[int] {[]int{1, 2, 3, 4, 5}, 0, 0, 5, []int{0, 0, 0, 0, 0}},
	fillTest[int] {[]int{1, 2, 3, 4, 5}, 0, 1, 4, []int{1, 0, 0, 0, 5}},
	fillTest[int] {[]int{1, 2, 3, 4, 5}, 0, 2, 3, []int{1, 2, 0, 4, 5}},
}

var stringFillTests = []fillTest[string] {
	fillTest[string] {[]string{"a", "b", "c", "d", "e"}, "", 0, 5, []string{"", "", "", "", ""}},
	fillTest[string] {[]string{"a", "b", "c", "d", "e"}, "", 1, 4, []string{"a", "", "", "", "e"}},
	fillTest[string] {[]string{"a", "b", "c", "d", "e"}, "", 2, 3, []string{"a", "b", "", "d", "e"}},
}

func TestFill(t *testing.T) {
	for idx, test := range intFillTests {
		actual := tools.Fill(test.slice, test.value, test.start, test.end)
		if tools.ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. Fill() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}

	for idx, test := range stringFillTests {
		actual := tools.Fill(test.slice, test.value, test.start, test.end)
		if tools.ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. Fill() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}
}