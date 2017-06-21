// +build !windows

package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func echoServer(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		log.Printf("Server got: %s", string(data))
		_, err = c.Write(data)
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}

func main() {
	l, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatalf("listen error: %s", err)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(l net.Listener, c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		l.Close()
		os.Exit(0)
	}(l, sigc)

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatalf("accept error: %s", err)
		}

		go echoServer(fd)
	}
}
