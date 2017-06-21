package main

import (
	"flag"
	"log"
	"runtime"
	"time"
)

func main() {
	var foo string
	flag.StringVar(&foo, "bar", "ham", "eggs")
	var bar time.Duration
	flag.DurationVar(&bar, "have", time.Second*60, "it")
	flag.Parse()
	//log.Printf("%v %s\n", foo, *foo)
	log.Printf("%+v %+v", foo, bar)

	//s := os.Getenv("GOMAXPROCS")
	//log.Print(s)

	i := runtime.NumCPU()
	log.Print(i)
	//log.Printf("%s %s %d\n", runtime.GOOS, runtime.GOARCH, runtime.GOMAXPROCS)
	log.Printf("%s %s\n", runtime.GOOS, runtime.GOARCH)

	runtime.Gosched()

	//d := time.Duration(time.Second * 2)
	//time.Sleep(d)
	log.Printf("%s %+v", bar, bar)
	//time.Sleep(bar)
}
