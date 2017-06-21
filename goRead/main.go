package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func foo() {
	f, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", buf)
}

func main() {
	foo()
	os.Exit(1)

	f, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := make([]byte, 100)
	_, err = f.Read(buf)
	for ; err == nil; _, err = f.Read(buf) {
		fmt.Printf("%s", buf)
		//fmt.Print(string(buf))
		buf = make([]byte, 100)
	}
	//fmt.Printf("%d: %s\n", n, buf)
	fmt.Printf("%s", buf)

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
}
