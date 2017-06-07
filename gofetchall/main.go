package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"text/tabwriter"
	"time"
)

func main() {
	ch := make(chan string)

	us := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8080",
		"http://foobar:8080",
	}

	for _, u := range us {
		go fetch(u, ch)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for i := 0; i < len(us); i++ {
		fmt.Fprintf(tw, "%s\n", <-ch)
	}
	tw.Flush()
}

func fetch(u string, ch chan string) {
	var res string

	start := time.Now()
	resp, err := http.Get(u)
	if err == nil {
		defer resp.Body.Close()
		n, errCopy := io.Copy(ioutil.Discard, resp.Body)
		if errCopy == nil {
			res = fmt.Sprintf("%d bytes\t%s duration", n, time.Since(start))
		} else {
			res = errCopy.Error()
		}

		/*
			body, errRead := ioutil.ReadAll(resp.Body)
			if errRead == nil {
				res = fmt.Sprintf("%s", body)
			} else {
				res = errRead.Error()
			}
		*/
	} else {
		res = err.Error()
	}

	ch <- fmt.Sprintf("%s\t%s", u, res)
}

/*
io.Copy
ioutil.Discard
time.Now()
time.Since()
http.Get()
resp.Body.Close() - Reader
*/
