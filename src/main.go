package main

import (
	"fmt"
	"net"
	"time"
)

var finished chan struct{}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func main() {
	listener, _ := net.Listen("tcp", "localhost:8080")
	conn, err := listener.Accept()

	go spinner(100 * time.Microsecond)
	fmt.Println("main goroutine")
}
