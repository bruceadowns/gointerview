/*
Execution instructions:

$ go run main.go
or
$ go run --race main.go
*/

package main

import (
	"log"
	"os"
	"strconv"
)

func fizzbuzz(iter int) {
	for i := 1; i <= iter; i++ {
		if i%15 == 0 {
			log.Print("FizzBuzz")
		} else if i%5 == 0 {
			log.Print("Buzz")
		} else if i%3 == 0 {
			log.Print("Fizz")
		} else {
			log.Print(i)
		}
	}
}

func sumfizzbuzz(iter int) {
	sum := 0
	for i := 1; i < iter; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	log.Print(sum)
}

func main() {
	iter := 10
	var err error
	if len(os.Args) > 1 {
		iter, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	//log.Printf("%d", iter)

	fizzbuzz(iter)
	sumfizzbuzz(iter)
}
