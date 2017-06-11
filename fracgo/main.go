package main

// #cgo LDFLAGS: -L . -lsum
// #include "sum.h"
import "C"

func main() {
	C.sum(4, 5)
}

// gcc -c -o sum.o sum.c
// ar -rsc libsum.a sum.o
