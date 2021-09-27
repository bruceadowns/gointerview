package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var tt = []struct {
		S string
		R int
	}{
		{"", -1},
		{"a", 0},
		{"b", 0},
		{"ab", 0},
		{"abab", 0},
		{"foobar", -1},
		{"babaa", 3},
		{"bbbab", 4},
		{"bbbaaabbb", 0},
		{"aabbaaabb", 3},
	}
	for _, test := range tt {
		r := solution(test.S)
		if r != test.R {
			t.Errorf("%s expected %d actual %d", test.S, test.R, r)
		}
	}
}
