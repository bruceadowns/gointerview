package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lineCount(r io.Reader) (res int) {
	for input := bufio.NewScanner(r); input.Scan(); {
		res++
	}

	return
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("stdin: %d\n", lineCount(os.Stdin))
	} else {
		for _, v := range os.Args[1:] {
			f, err := os.Open(v)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			defer f.Close()

			fmt.Printf("%s: %d\n", v, lineCount(f))
		}
	}
}
