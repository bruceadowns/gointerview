/*
You are given a string S consisting of letters 'a' and/or 'b'. A block is a consecutive fragment of S composed of the same letters and surrounded by different letters or string endings. For example, S = "abbabbaaa" has five blocks: "a", "bb", "a", "bb" and "aaa".

What is the minimum number of additional letters needed to obtain a string containing blocks of equal lengths? Letters can be added at the beginning, between two existing letters, or at the end of the string.

Write a function:

class Solution { public int solution(String S); }

that, given a string S of length N, returns the minimum number of additional letters needed to obtain a string containing blocks of equal lengths.

Examples:

1. Given S = "babaa", the function should return 3. There are four blocks: "b", "a", "b", "aa". One letter each should be added to the first, second and third blocks, therefore obtaining a string "bbaabbaa", in which every block is of equal length.

2. Given S = "bbbab", the function should return 4. Two letters each should be added to the second and third blocks, therefore obtaining a string "bbbaaabbb", in which every block is of equal length.

3. Given S = "bbbaaabbb", the function should return 0. All blocks are already of equal lengths.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..40,000];
string S consists only of the characters "a" and/or "b".
*/

/*
Execution instructions:

$ go run main.go
or
$ go run --race main.go
*/

package main

import (
	"log"
	"math"
	"os"
)

const (
	Unknown = iota
	A
	B
)

func solution(S string) int {
	if len(S) < 1 || len(S) > 40000 {
		return -1
	}

	var buckets []int
	max := math.MinInt32
	curr := 0
	state := Unknown
	for _, s := range S {
		switch s {
		case 'a':
			switch state {
			case Unknown:
				state = A
				curr++
			case A:
				curr++
			case B:
				if curr > max {
					max = curr
				}
				buckets = append(buckets, curr)
				curr = 1
				state = A
			}

		case 'b':
			switch state {
			case Unknown:
				state = B
				curr++
			case A:
				if curr > max {
					max = curr
				}
				buckets = append(buckets, curr)
				curr = 1
				state = B
			case B:
				curr++
			}

		default:
			return -1
		}
	}
	if curr > max {
		max = curr
	}
	buckets = append(buckets, curr)

	//log.Printf("max %d", max)
	//log.Printf("buckets %v", buckets)

	ret := 0
	for _, i := range buckets {
		ret += max - i
	}
	return ret
}

func main() {
	S := "babaa"
	if len(os.Args) == 2 {
		S = os.Args[1]
	}

	log.Printf("solution for '%s' is %d", S, solution(S))
}
