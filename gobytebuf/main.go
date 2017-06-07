package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.Buffer{}

	buf.WriteString("foo")
	for i := 0; i < 100; i++ {
		buf.WriteString(fmt.Sprintf("%d", i))
	}
	buf.WriteString("bar")

	fmt.Println(fmt.Sprintf("%s\n", buf.String()))
}
