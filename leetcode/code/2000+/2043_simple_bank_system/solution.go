package _043_simple_bank_system

type Bank struct {
	balance []int64
	total   int
}

func Constructor(balance []int64) Bank {
	return Bank{
		balance: balance,
		total:   len(balance),
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if this.illegalAccount(account1) || this.illegalAccount(account2) || this.balance[account1-1] < money {
		return false
	}
	this.balance[account1-1] -= money
	this.balance[account2-1] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if this.illegalAccount(account) {
		return false
	}
	this.balance[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if this.illegalAccount(account) || this.balance[account-1] < money {
		return false
	}
	this.balance[account-1] -= money
	return true
}

func (this *Bank) illegalAccount(account int) bool {
	if account < 0 || account > this.total {
		return true
	}
	return false
}
