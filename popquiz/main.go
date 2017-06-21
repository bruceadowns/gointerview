package main

import "fmt"

func main() {
	x := make([]string, 0, 3)
	fmt.Println(x)

	x = append(x, "a", "b")
	fmt.Println(x)

	y := append(x, "c")
	fmt.Println(y)

	_ = append(x, "d")
	fmt.Println(y)
}
