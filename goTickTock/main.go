package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func tick(ctx context.Context, wg *sync.WaitGroup) {
	// tell the caller we've stopped
	defer wg.Done()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case now := <-ticker.C:
			fmt.Printf("tick: tick %s\n", now.UTC().Format("20060102-150405.000000000"))
		case <-ctx.Done():
			fmt.Println("tick: caller has told us to stop")
			return
		}
	}
}

func tock(ctx context.Context, wg *sync.WaitGroup) {
	// tell the caller we've stopped
	defer wg.Done()

	ticker := time.NewTicker(4 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case now := <-ticker.C:
			fmt.Printf("tock: tock %s\n", now.UTC().Format("20060102-150405.000000000"))
		case <-ctx.Done():
			fmt.Println("tock: caller has told us to stop")
			return
		}
	}
}

func main() {
	//_ = time.NewTicker(-1)

	fmt.Println("main: starting")

	// create a context that we can cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// a WaitGroup for the goroutines to tell us they've stopped
	wg := sync.WaitGroup{}

	// a channel for `tick()` to tell us they've stopped
	wg.Add(1)
	go tick(ctx, &wg)

	// a channel for `tock()` to tell us they've stopped
	wg.Add(1)
	go tock(ctx, &wg)

	// listen for C-c
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("main: received C-c - shutting down")

	// tell the goroutines to stop
	fmt.Println("main: telling goroutines to stop")
	cancel()

	// and wait for them both to reply back
	wg.Wait()
	fmt.Println("main: all goroutines have told us they've finished")
}
