package main

//
// https://joeshaw.org/dont-defer-close-on-writable-files/
// close could fail, if bytes have not been sync'ed
//

import (
	"log"
	"os"
)

func idiomaticWrite(fileName, content string) (int, error) {
	//
	// idiomatic case
	// close could fail if bytes are not flushed
	// thus losing data and error condition
	//

	f, err := os.Create(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	n, err := f.WriteString(content)
	if err != nil {
		return 0, nil
	}

	return n, err
}

func doubleCloseWrite1(fileName, content string) (n int, err error) {
	//
	// defer and explicitly close
	// not ideal
	//

	f, err := os.Create(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	n, err = f.WriteString(content)
	if err != nil {
		return 0, err
	}

	err = f.Close()
	if err != nil {
		return 0, err
	}

	return n, nil
}

func doubleCloseWrite2(fileName, content string) (int, error) {
	//
	// attempt close in all paths
	// not ideal
	//

	f, err := os.Create(fileName)
	if err != nil {
		return 0, err
	}

	n, err := f.WriteString(content)
	if err != nil {
		f.Close()
		return 0, nil
	}

	return n, f.Close()
}

func deferCloseFuncWrite(fileName, content string) (n int, err error) {
	//
	// defer closure which magically sets err
	// not ideal
	//

	f, err := os.Create(fileName)
	if err != nil {
		return 0, err
	}

	defer func() {
		cerr := f.Close()
		if cerr != nil {
			err = cerr
		}
	}()

	n, err = f.WriteString(content)
	return
}

func closeSyncWrite(fileName, content string) (int, error) {
	//
	// defer close, and use sync to ensure flush
	//

	f, err := os.Create(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	n, err := f.WriteString(content)
	if err != nil {
		return 0, nil
	}

	return n, f.Sync()
}

func main() {
	fileName := "file.txt"
	contents := "foobar"

	// fn := idiomaticWrite
	// fn := doubleCloseWrite1
	// fn := doubleCloseWrite2
	// fn := deferCloseFuncWrite
	fn := closeSyncWrite

	n, err := fn(fileName, contents)
	if err != nil {
		log.Fatalf("Error writing an atom: %s", err)
	}
	log.Printf("Wrote %d chars", n)
}
