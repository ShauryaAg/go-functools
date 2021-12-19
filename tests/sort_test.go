package tests

import (
	"testing"

	tools "tools/tools"
)

type sortTest[T any] struct {
	slice []T
	compareFn func(T, T) bool
	expected []T
}

var intSortTests = []sortTest[int] {
	sortTest[int] {[]int{1, 3, 4, 5, 8, 10, 2}, tools.DefaultCompare[int], []int{1, 2, 3, 4, 5, 8, 10}},
	sortTest[int] {[]int{1, 3, 4, 5, 8, 10, 2}, func(a, b int) bool { return a > b }, []int{1, 2, 3, 4, 5, 8, 10}},
	sortTest[int] {[]int{1, 3, 4, 5, 8, 10, 2}, func(a, b int) bool { return a < b }, []int{10, 8, 5, 4, 3, 2, 1}},
}

var stringSortTests = []sortTest[string] {
	sortTest[string] {[]string{"a", "c", "d", "e", "g", "k", "b"}, tools.DefaultCompare[string], []string{"a", "b", "c", "d", "e", "g", "k"}},
	sortTest[string] {[]string{"a", "c", "d", "e", "g", "k", "b"}, func(a, b string) bool { return a > b }, []string{"a", "b", "c", "d", "e", "g", "k"}},
	sortTest[string] {[]string{"a", "c", "d", "e", "g", "k", "b"}, func(a, b string) bool { return a < b }, []string{"k", "g", "e", "d", "c", "b", "a"}},
}

func TestSort(t *testing.T) {
	for idx, test := range intSortTests {
		actual := tools.Sort(test.slice, test.compareFn)
		if tools.ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. Sort() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}

	for idx, test := range stringSortTests {
		actual := tools.Sort(test.slice, test.compareFn)
		if tools.ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. Sort() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}
}