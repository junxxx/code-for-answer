package channel

import (
	"fmt"
	"time"
)

var c = make(chan int)

func Send(timers int) {
	for {
		if timers == 0 {
			break
		}
		fmt.Printf("There is left %d times send\n", timers)
		c <- timers
		timers--
		time.Sleep(1 * time.Second)
	}
	close(c)
}

func Run() {
	go Send(5)

	for {
		if i, ok := <-c; ok {
			fmt.Println("receive from channel", i)
		} else {
			break
		}
	}
	fmt.Println("finished receive from channel")
}
