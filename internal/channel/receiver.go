package channel

import (
	"fmt"
	"time"
)

var done = make(chan struct{})

func receive(num int) {
	for {
		if cancelled() {
			return
		}
		fmt.Println("num receiving ", num, cancelled())
		fmt.Println("num ", num, <-done, cancelled())
	}
}

func trigger() {
	fmt.Println("ready for closing channel")
	time.Sleep(3 * time.Second)
	fmt.Println("close channel")
	close(done)
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func Entrypoint() {
	go trigger()
	go receive(1)
	go receive(2)

	time.Sleep(20 * time.Second)
}
