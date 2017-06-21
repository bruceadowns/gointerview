package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"time"
)

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		log.Println(msg)
	case <-ctx.Done():
		log.Printf("%s", ctx.Err())
	}

	/*
		select {
		case <-time.After(d):
			fmt.Println(msg)
		}
	*/

	//time.Sleep(d)
	//fmt.Println(msg)
}

func main() {
	// root
	ctx := context.Background()

	//ctx, cancel := context.WithTimeout(ctx, time.Second)
	//defer cancel()

	ctx, cancel := context.WithCancel(ctx)
	//time.AfterFunc(2*time.Second, cancel)

	go func() {
		//time.Sleep(1 * time.Second)
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	sleepAndTalk(ctx, 3*time.Second, "Hello")
}

// pre-1.7
// x/net/context

// create context
// receive context
// http client / server with context

// cancellation/propogation usage per package
// send values for logging
// CancelFunc
// Context type interface
// Context.Background
// no cancelation or timeout
// root context and children/tree

// Background
// WithCancel
// WithTimeout
// WithDeadline
