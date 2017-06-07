package main

import (
	"fmt"
	"sync"
)

type list []int
type listmap map[int]list

const (
	maxlen = 10
)

func print(id int, l []int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for _, v := range l {
		fmt.Printf("%d\t%d\n", id, v)
	}
}

func main() {
	fmt.Println("Start")

	// make list
	l := make(listmap)
	for i := 0; i < maxlen; i++ {
		childlist := make(list, maxlen)
		for j := 0; j < maxlen; j++ {
			childlist[j] = (j + 1) * (i + 1)
		}
		l[i] = childlist
	}

	// print list in parallel
	var wg sync.WaitGroup
	for k, v := range l {
		go print(k, v, &wg)
	}
	wg.Wait()

	fmt.Println("Done")
}
