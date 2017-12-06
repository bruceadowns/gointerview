package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type node struct {
	on      bool
	visited bool
}
type row []node
type matrix []row

func input(in io.Reader) (m matrix) {
	prevLineLen := 0
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		if prevLineLen != 0 && prevLineLen != len(fields) {
			log.Fatalf("invalid input %d != %d", prevLineLen, len(fields))
		}
		prevLineLen = len(fields)

		r := make(row, 0)
		for _, field := range fields {
			switch field {
			case "0":
				r = append(r, node{})
			case "1":
				r = append(r, node{on: true})
			default:
				log.Fatalf("invalid input: %s", field)
			}
		}

		m = append(m, r)
	}

	return
}

func visit(m matrix, x int, y int) {
	m[x][y].visited = true

	for i := x - 1; i < x+2; i++ {
		if i < 0 || i > len(m)-1 {
			continue
		}

		for j := y - 1; j < y+2; j++ {
			if j < 0 || j > len(m[i])-1 {
				continue
			}

			if !m[i][j].visited && m[i][j].on {
				visit(m, i, j)
			}
		}
	}
}

func main() {
	m := input(os.Stdin)

	count := 0
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[x]); y++ {
			if !m[x][y].visited && m[x][y].on {
				visit(m, x, y)
				count++
			}
		}
	}

	log.Printf("island count: %d", count)
}
