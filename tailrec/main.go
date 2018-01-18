package main

import "log"

func factTR(n, a int) int {
	if n == 0 {
		return a
	}

	return factTR(n-1, n*a)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	{
		val := fact(10)
		log.Print(val)
	}

	{
		val := factTR(10, 1)
		log.Print(val)
	}
}
