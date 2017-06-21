package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func doSomethingAwesome(ctx context.Context) {
	fmt.Println("Sleep")
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout")
	case <-ctx.Done():
		fmt.Println("Done")
	}
}

func main() {
	ctx := context.Background()

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		fmt.Println("Stop")
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			fmt.Println("cancel")
			cancel()
		case <-ctx.Done():
		}
	}()

	doSomethingAwesome(ctx)
}
