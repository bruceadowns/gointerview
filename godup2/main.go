package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"text/tabwriter"
)

type fileCount map[string]int
type fileCounts map[string]fileCount

func count(r io.Reader) (fc fileCount) {
	fc = make(map[string]int)
	input := bufio.NewScanner(r)
	for input.Scan() {
		if len(input.Text()) > 0 {
			fc[input.Text()]++
		}
	}
	return fc
}

func main() {
	counts := make(fileCounts)
	if len(os.Args) == 1 {
		counts["stdin"] = count(os.Stdin)
	} else {
		for _, arg := range os.Args[1:] {
			if f, err := os.Open(arg); err == nil {
				counts[arg] = count(f)
			} else {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for k, v := range counts {
		fmt.Fprintf(w, "%s\n", k)
		for l, c := range v {
			fmt.Fprintf(w, fmt.Sprintf("\t%v\t%d\n", l, c))
		}
	}
	w.Flush()
}
