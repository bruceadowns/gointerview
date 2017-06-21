package main

import (
	"io"
	"log"
	"net"
	"time"
)

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		nbytes, err := r.Read(buf[:])
		if err != nil {
			return
		}
		log.Printf("Client got [%d] %s\n", nbytes, string(buf[0:nbytes]))
	}
}

func main() {
	c, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	go reader(c)
	for {
		nbytes, err := c.Write([]byte("hi"))
		if err != nil {
			log.Fatal("write error:", err)
			break
		}
		log.Printf("wrote %d bytes\n", nbytes)
		time.Sleep(time.Second)
	}
}
