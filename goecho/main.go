package main

import (
	"fmt"
	"os"
)

func main() {
	for k, v := range os.Args {
		sep := ""
		if k < len(os.Args)-1 {
			sep = " "
		}
		fmt.Printf("%s%s", v, sep)
	}

	fmt.Println()
}
