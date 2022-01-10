package bank

// It's not a concurrency-safe implementation

var balance int

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}
