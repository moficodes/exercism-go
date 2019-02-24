package account

import "sync"

// Open(initialDeposit int64) *Account
// (*Account) Close() (payout int64, ok bool)
// (*Account) Balance() (balance int64, ok bool)
// (*Account) Deposit(amount int64) (newBalance int64, ok bool)

// Account represents a bank account
type Account struct {
	// mu protects both variable
	mu      sync.Mutex
	balance int64
	open    bool
}

// Open a new accout with the initial ammount
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{
		balance: initialDeposit,
		open:    true,
		mu:      sync.Mutex{},
	}
}

// Close an existing account
func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.open {
		return 0, false
	}
	payout := a.balance
	a.open = false
	a.balance = 0
	return payout, true
}

// Balance return the current balance
func (a *Account) Balance() (int64, bool) {
	if !a.open {
		return 0, false
	}
	return a.balance, true
}

// Deposit add money to our account
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.open || a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
