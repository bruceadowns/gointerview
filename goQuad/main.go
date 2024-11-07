package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func getFloat64() (res float64) {
	scanner.Scan()
	line := scanner.Text()

	var err error
	res, err = strconv.ParseFloat(line, 10)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func calcQuad(a, b, c float64) (float64, float64) {
	d := math.Pow(b, 2) - 4*a*c
	x := (-b + math.Sqrt(d)) / (2 * a)
	y := (-b - math.Sqrt(d)) / (2 * a)

	return x, y
}

func main() {
	fmt.Printf("What is a? ")
	a := getFloat64()

	fmt.Printf("What is b? ")
	b := getFloat64()

	fmt.Printf("What is c? ")
	c := getFloat64()

	x, y := calcQuad(a, b, c)
	fmt.Printf("x: %.2f y: %.2f\n", x, y)
}
