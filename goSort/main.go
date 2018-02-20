package main

import "log"

func insertionSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

func main() {
	in := []int{1, 3, 5, 2, 4, 0, 6}
	log.Print(in)

	insertionSort(in)
	log.Print(in)
}
