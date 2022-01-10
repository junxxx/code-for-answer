package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/junxxx/codeforgetanswer/concurrency/bank"
)

func TestConcurrency(t *testing.T) {
	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
	}()

	// Bob
	go bank.Deposit(100)

	// waiting for goroutine finish
	time.Sleep(1 * time.Second)
}
