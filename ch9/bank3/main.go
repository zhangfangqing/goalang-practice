package bank

import "sync"

var (
	mu      sync.Mutex
	balance int
)

func Wiithdarw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func Deposit(amount int) {
	mu.Lock()
	deposit(amount)
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

func deposit(amount int) { balance += amount }
