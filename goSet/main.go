package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	set := make(map[int]struct{})

	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 101; i++ {
		j := rand.Intn(10)
		fmt.Println(j)

		if _, ok := set[j]; !ok {
			set[j] = struct{}{}
		}
	}

	for k := range set {
		fmt.Printf("%d ", k)
	}
	fmt.Println()
}
