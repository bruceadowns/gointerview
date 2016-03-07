package main

import (
	"math"
	"testing"
)

func printDupsCount(t *testing.T, a []int) {
	count := make([]bool, len(a))

	for _, v := range a {
		if count[v] {
			t.Logf("Dup %d", v)
		} else {
			count[v] = true
		}
	}
}

func printDupsMap(t *testing.T, a []int) {
	m := make(map[int]struct{})

	for _, v := range a {
		if _, ok := m[v]; ok {
			t.Logf("Dup %d", v)
		} else {
			m[v] = struct{}{}
		}
	}
}

func TestDups(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 1}
	printDupsCount(t, a)
	printDupsMap(t, a)
}

func maxDiff(prices []int) (diff int) {
	diff = math.MinInt32
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices)-1; j++ {
			if prices[j]-prices[i] > diff {
				diff = prices[j] - prices[i]
			}
		}
	}

	return
}

// this function is deficient
func maxProfit(prices []int) (buyIndex, sellIndex, profit int) {
	minPrice := math.MaxInt32
	for k, v := range prices {
		if v < minPrice {
			buyIndex = k
			minPrice = v
		}
	}

	maxPrice := 0
	for i := buyIndex; i < len(prices)-1; i++ {
		if prices[i] > maxPrice {
			sellIndex = i
			maxPrice = prices[i]
		}
	}

	profit = maxPrice - minPrice

	return
}

func TestStock(t *testing.T) {
	var tt = []struct {
		in []int
	}{
		{
			[]int{2, 3, 10, 6, 4, 8, 1},
		},
		{
			[]int{1, 2, 90, 10, 110},
		},
		{
			[]int{500, 499, 498, 497, 496, 495},
		},
	}

	for _, test := range tt {
		buy, sell, profit := maxProfit(test.in)
		t.Logf("buyIndex:%d sellIndex:%d profit:%d", buy, sell, profit)

		diff := maxDiff(test.in)
		t.Logf("diff:%d", diff)
	}
}

func count(s, r string) int {
	if len(s) == 0 {
		return 1
	}

	var res int
	for i := 0; i < len(r); i++ {
		if s[0] == r[i] {
			res += count(s[1:], r[i+1:])
		}
	}

	return res
}

func TestCount(t *testing.T) {
	var tt = []struct {
		s string
		r string
		e int
	}{
		{
			s: "bar",
			r: "barbar",
			e: 4,
		},
		{
			s: "abc",
			r: "abcdefab",
			e: 1,
		},
	}

	for _, test := range tt {
		c := count(test.s, test.r)
		if c != test.e {
			t.Fatalf("invalid count %d for %v", c, test)
		}
	}
}

type block int
type maze [][]block
type coord struct {
	x, y int
}
type solution []coord

func solveMazeR(m maze, c coord) (sol solution, solved bool) {
	m[c.x][c.y] = 1
	sol = append(sol, c)

	if c.x == len(m)-2 &&
		c.y == len(m[0])-2 {
		solved = true
		return
	}

	neighbors := []coord{
		{c.x - 1, c.y},
		{c.x, c.y - 1},
		{c.x, c.y + 1},
		{c.x + 1, c.y},
	}

	for _, n := range neighbors {
		if m[n.x][n.y] == 0 {
			var solR solution
			if solR, solved = solveMazeR(m, n); solved {
				sol = append(sol, solR...)
				return
			}
		}
	}

	return
}

func solveMaze(m maze) (s solution) {
	s, _ = solveMazeR(m, coord{1, 1})
	return s
}

func TestSolveMaze(t *testing.T) {
	m := maze{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 1, 1},
		{1, 0, 1, 1, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}

	s := solveMaze(m)
	t.Log(s)
}
