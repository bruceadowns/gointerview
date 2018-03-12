package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//
	// run via
	// cat test.txt | go run main.go
	// cat test.txt | go run main.go -once
	//

	var once bool
	flag.BoolVar(&once, "once", false, "one occurrence allowed")
	flag.Parse()

	// process each line from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("line: %s", line)

		// tokenize input into slice of ints
		is := make([]int, 0)
		for _, v := range strings.Fields(line) {
			i, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}

			is = append(is, i)
		}

		// hold map of int to its occurrence
		m := make(map[int]int)
		for _, v := range is {
			// insert or update occurrence
			m[v]++

			partner := 7 - v
			// check for seven partner
			if p, ok := m[partner]; ok && p > 0 {
				// found valid partner
				if once {
					// print line of partner occurrence
					log.Printf("%d + %d = %d", v, partner, 7)
					m[partner] = p - 1
				} else {
					// print line per partner occurrence
					for i := 0; i < p; i++ {
						log.Printf("%d + %d = %d", v, partner, 7)
					}
				}
			}
		}
	}
}
