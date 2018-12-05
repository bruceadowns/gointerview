package main

import (
	"log"
	"math"
	"strconv"
)

func isPrime(i int) bool {
	if i == 1 {
		return false
	}

	for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
		if i%j == 0 {
			return false
		}
	}

	return true
}

func main() {
	// unoptimized loop
outer:
	for i := 10000000; i < 1000000000; i++ {
		s := strconv.Itoa(i)
		for j := 0; j < 8; j++ {
			k, _ := strconv.Atoi(s[:len(s)-j])
			if !isPrime(k) {
				continue outer
			}
		}
		log.Print(i)
	}

	// 23399339
	// 29399999
	// 37337999
	// 59393339
	// 73939133
}
