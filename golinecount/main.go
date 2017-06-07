package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lineCount(r io.Reader) (res int) {
	for input := bufio.NewScanner(r); input.Scan(); {
		if len(input.Text()) > 0 {
			res++
		}
	}

	return
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("stdin: %d\n", lineCount(os.Stdin))
	} else {
		for _, v := range os.Args[1:] {
			if f, err := os.Open(v); err == nil {
				fmt.Printf("%s: %d\n", v, lineCount(f))
			} else {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}
