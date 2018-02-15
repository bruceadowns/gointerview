package main

import (
	"log"
)

/*
Given a collection of intervals, merge all overlapping intervals.
For example,
Given [1,3],[2,6],[8,10],[15,18],
return [1,6],[8,10],[15,18].
*/

type interval struct {
	begin, end int
}

func min(i, j int) int {
	if i <= j {
		return i
	}

	return j
}

func max(i, j int) int {
	if i >= j {
		return i
	}

	return j
}

func (i interval) overlaps(j interval) bool {
	if i.end >= j.begin {
		return true
	}

	return false
}

func (i interval) merge(j interval) interval {
	return interval{begin: min(i.begin, j.begin), end: max(i.end, j.end)}
}

func main() {
	input := []interval{
		{begin: 1, end: 3},
		{begin: 2, end: 6},
		{begin: 8, end: 10},
		{begin: 15, end: 18},
		{begin: 18, end: 19},
	}

	prev := input[0]
	ol := make([]interval, 0)
	for i := 1; i < len(input); i++ {
		if prev.overlaps(input[i]) {
			prev = prev.merge(input[i])
		} else {
			ol = append(ol, prev)
			prev = input[i]
		}
	}
	ol = append(ol, prev)

	for _, v := range ol {
		log.Print(v)
	}
}
