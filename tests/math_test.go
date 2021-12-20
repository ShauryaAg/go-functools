package tests

import (
	"testing"
	"math"

	utils "github.com/ShauryaAg/go-functools/utils"
)

type mathTest[T any] struct {
	a T
	b T
	expected T
}

var intMinTest = []mathTest[int] {
	mathTest[int] {1, 2, 1},
	mathTest[int] {2, 1, 1},
	mathTest[int] {1, 1, 1},
	mathTest[int] {1, int(math.Inf(-1)), int(math.Inf(-1))},
}

var intMaxTest = []mathTest[uint64] {
	mathTest[uint64] {1, 2, 2},
	mathTest[uint64] {2, 1, 2},
	mathTest[uint64] {1, 1, 1},
	mathTest[uint64] {uint64(1), uint64(math.Inf(1)), uint64(math.Inf(1))},
}

var stringMinTest = []mathTest[string] {
	mathTest[string] {"a", "b", "a"},
	mathTest[string] {"b", "a", "a"},
	mathTest[string] {"a", "a", "a"},
}

var stringMaxTest = []mathTest[string] {
	mathTest[string] {"a", "b", "b"},
	mathTest[string] {"b", "a", "b"},
	mathTest[string] {"a", "a", "a"},
}

func TestMin(t *testing.T) {
	for idx, test := range intMinTest {
		if utils.Min(test.a, test.b) != test.expected {
			t.Errorf("%d. Min(%d, %d) = %d, expected %d", idx, test.a, test.b, utils.Min(test.a, test.b), test.expected)
		}
	}

	for idx, test := range stringMinTest {
		if utils.Min(test.a, test.b) != test.expected {
			t.Errorf("%d. Min(%s, %s) = %s, expected %s", idx, test.a, test.b, utils.Min(test.a, test.b), test.expected)
		}
	}
}

func TestMax(t *testing.T) {
	for idx, test := range intMaxTest {
		if utils.Max(test.a, test.b) != test.expected {
			t.Errorf("%d. Max(%d, %d) = %d, expected %d", idx, test.a, test.b, utils.Max(test.a, test.b), test.expected)
		}
	}

	for idx, test := range stringMaxTest {
		if utils.Max(test.a, test.b) != test.expected {
			t.Errorf("%d. Max(%s, %s) = %s, expected %s", idx, test.a, test.b, utils.Max(test.a, test.b), test.expected)
		}
	}
}
