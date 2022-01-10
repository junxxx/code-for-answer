// a concurrency-safe with one account
package safebank

import "sync"

var (
	mu      sync.RWMutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}
func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}
