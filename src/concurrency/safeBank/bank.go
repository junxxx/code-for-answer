// a concurrency-safe with one account
package safebank

import "sync"

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}
func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}
