package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/junxxx/codeforgetanswer/concurrency/bank"
	safebank "github.com/junxxx/codeforgetanswer/concurrency/safeBank"
)

func testConcurrency(t *testing.T) {
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

func testSafeBank(t *testing.T) {
	var n sync.WaitGroup
	// Alice
	n.Add(1)
	go func() {
		safebank.Deposit(200)
		fmt.Println("safe version balance =", safebank.Balance())
		n.Done()
	}()

	// Bob
	n.Add(1)
	go func() {
		safebank.Deposit(100)
		n.Done()
	}()
	// waiting for goroutine finish
	n.Wait()
}
