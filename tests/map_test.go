package tests

import (
	"testing"

	tools "github.com/ShauryaAg/go-functools/tools"
)

type mapTest[T any, U any] struct {
	slice    []T
	callback func(T) U
	expected []U
}

var intMapTest = []mapTest[int, int]{
	mapTest[int, int]{[]int{1, 2, 3, 4, 5}, func(val int) int { return val * 2 }, []int{2, 4, 6, 8, 10}},
	mapTest[int, int]{[]int{1, 2, 3, 4, 5}, func(val int) int { return val + 1 }, []int{2, 3, 4, 5, 6}},
	mapTest[int, int]{[]int{1, 2, 3, 4, 5}, func(val int) int { return val * val }, []int{1, 4, 9, 16, 25}},
}

var stringMapTest = []mapTest[string, string]{
	mapTest[string, string]{[]string{"a", "b", "c", "d", "e"}, func(val string) string { return val + "!" }, []string{"a!", "b!", "c!", "d!", "e!"}},
}

var runeIntMapTest = []mapTest[rune, int]{
	mapTest[rune, int]{[]rune{'a', 'b', 'c', 'd', 'e'}, func(val rune) int { return int(val) }, []int{97, 98, 99, 100, 101}},
}

func TestMap(t *testing.T) {
	for idx, test := range intMapTest {
		actual := tools.Map(test.slice, test.callback)
		if tools.ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. Map() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}

	for idx, test := range stringMapTest {
		actual := tools.Map(test.slice, test.callback)
		if tools.ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. Map() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}

	for idx, test := range runeIntMapTest {
		actual := tools.Map(test.slice, test.callback)
		if tools.ArrayEquals(actual, test.expected) == false {
			t.Errorf("%d. Map() returned %v, expected %v", idx+1, actual, test.expected)
		}
	}
}
