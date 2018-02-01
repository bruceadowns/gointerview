package main

import "log"

func binarySearch(sortedList []int, lookingFor int) int {
	lo := 0
	hi := len(sortedList) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		midValue := sortedList[mid]
		log.Println("Middle value is:", midValue)

		if midValue == lookingFor {
			return mid
		} else if midValue > lookingFor {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return -1
}

func main() {
	lookingFor := 6
	sortedList := []int{1, 3, 4, 6, 7, 9, 10, 11, 13}
	log.Println("Looking for", lookingFor, "in the sorted list:", sortedList)

	index := binarySearch(sortedList, lookingFor)
	if index >= 0 {
		log.Println("Found the number", lookingFor, "at:", index)
	} else {
		log.Println("Did not find the number", lookingFor)
	}
}
