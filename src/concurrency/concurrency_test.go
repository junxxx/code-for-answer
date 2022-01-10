package main

import (
	"fmt"
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

func TestSafeBank(t *testing.T) {
	// Alice
	go func() {
		safebank.Deposit(200)
		fmt.Println("safe version balance =", safebank.Balance())
	}()

	// Bob
	go safebank.Deposit(100)

	// waiting for goroutine finish
	time.Sleep(1 * time.Second)
}
