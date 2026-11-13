package bankaccount

import "sync"

type Account struct {
	mu      sync.Mutex
	balance int64
	open    bool
	close   bool
}

func Open(amt int64) *Account {
	if amt < 0 {
		return nil
	}
	return &Account{mu: sync.Mutex{}, balance: amt, open: true, close: false}
}

func (a *Account) Balance() (bal int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.open || a.close {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amt int64) (newBal int64, ok bool) {
	if !a.open || a.close {
		return 0, false
	}

	if amt <= 0 {
		return a.Withdraw(-amt)
	}

	a.mu.Lock()
	defer a.mu.Unlock()
	if newBalance := a.balance + amt; newBalance < 0 {
		return a.balance, false
	} else {
		a.balance = newBalance
		return newBalance, true
	}
}

func (a *Account) Withdraw(amt int64) (newBal int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.open || a.close {
		return 0, false
	}

	if amt <= 0 {
		return a.balance, false
	}

	if newBalance := a.balance - amt; newBalance < 0 {
		return a.balance, false
	} else {
		a.balance = newBalance
		return newBalance, true
	}
}

func (a *Account) Close() (pay int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.open || a.close {
		return 0, false
	}

	pay = a.balance
	a.close = true
	return pay, true
}
