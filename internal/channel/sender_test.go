package channel

import "testing"

func TestSender(t *testing.T) {
	Run()
}

func TestConsumer(t *testing.T) {
	times := 10
	go Producer(times)

	for i := 0; i < 2*times; i++ {
		go Consumer(i)
	}
	Consumer(-1)
}
