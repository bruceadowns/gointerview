package main

import "log"

func main() {

	a := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	b := []int{2, 4, 5, 6, 8, 10, 12, 14, 17, 18, 20}

	{
		count := 0
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(b); j++ {
				if a[i] == b[j] {
					count++
				}
			}
		}
		log.Printf("count: %d O(n^2)", count)
	}

	{
		count := 0
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(b); j++ {
				if a[i] == b[j] {
					count++
					break
				}
			}
		}
		log.Printf("count: %d O(n^2)", count)
	}

	{
		count := 0
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(b); j++ {
				if a[i] == b[j] {
					count++
					break
				} else if a[i] < b[j] {
					break
				}
			}
		}
		log.Printf("count: %d O(n log n)", count)
	}
}
