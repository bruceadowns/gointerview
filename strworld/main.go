package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// #include <stdlib.h>
import "C"

type office int

const (
	boston office = iota
	portland
)

var officePlace = make(map[office]string)

var blah = []string{
	"foo",
	"bar",
}

type world struct{}

func (o office) String() string {
	return officePlace[o]
}

func (w world) String() string {
	return "World"
}

func init() {
	officePlace[boston] = "Boston"
	officePlace[portland] = "Portland"
}

func myPrint(args ...interface{}) {
	for _, arg := range args {
		switch v := reflect.ValueOf(arg); v.Kind() {
		case reflect.String:
			os.Stdout.WriteString(v.String())
		case reflect.Int:
			os.Stdout.WriteString(strconv.FormatInt(v.Int(), 10))
		default:
			panic("unknown type")
		}
	}
}

func main() {
	var w world
	fmt.Printf("Hello %s\n", w)
	fmt.Printf("Hello %s\n", new(world))
	fmt.Printf("Hello %s\n", &world{})

	fmt.Printf("Hello %s\n", boston)
	fmt.Printf("Hello %s\n", portland)

	/*
	start := time.Now()
	time.Sleep(time.Second * 1)
	duration := time.Since(start)
	fmt.Printf("duration: %s (%d)\n", duration, duration)
	*/

	myPrint("foo", 42, "\n")

	h := hex.Dumper(os.Stdout)
	//defer h.Close()
	fmt.Fprintf(h, "Hello World")
	h.Close()

	fmt.Printf("c rand: %d\n", C.random())
}
