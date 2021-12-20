package tests

import (
	"testing"

	tools "github.com/ShauryaAg/go-functools/tools"
)

type joinTest[T any] struct {
	slice     []T
	separator string
	expected  string
}

var intJoinTest = []joinTest[int]{
	joinTest[int]{[]int{1, 2, 3, 4, 5}, ",", "1,2,3,4,5"},
	joinTest[int]{[]int{1, 2, 3, 4, 5}, "", "12345"},
}

var stringJoinTest = []joinTest[string]{
	joinTest[string]{[]string{"a", "b", "c", "d", "e"}, ",", "a,b,c,d,e"},
	joinTest[string]{[]string{"a", "b", "c", "d", "e"}, "", "abcde"},
}

func TestJoin(t *testing.T) {
	for idx, test := range intJoinTest {
		actual := tools.Join(test.slice, test.separator)
		if actual != test.expected {
			t.Errorf("%d. Join() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}

	for idx, test := range stringJoinTest {
		actual := tools.Join(test.slice, test.separator)
		if actual != test.expected {
			t.Errorf("%d. Join() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}
}
