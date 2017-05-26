package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//
	// run via
	// cat test.txt | go run main.go
	//

	// process each line from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line: %s\n", line)

		// tokenize input into slice of ints
		is := make([]int, 0)
		for _, v := range strings.Fields(line) {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			is = append(is, i)
		}

		// hold map of int to its occurence
		m := make(map[int]int)
		for _, v := range is {
			// insert or update occurence
			if i, ok := m[v]; ok {
				m[v] = i + 1
			} else {
				m[v] = 1
			}

			// check for seven partner
			if p, ok := m[7-v]; ok {
				// found partner
				// print line per partner occurence
				for i := 0; i < p; i++ {
					fmt.Printf("%d + %d = 7\n", v, 7-v)
				}
			}
		}

		fmt.Println()
	}
}
