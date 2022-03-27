package pro2043

type Bank struct {
	balance map[int]int64
}

func Constructor(balance []int64) Bank {
	b := Bank{
		balance: make(map[int]int64),
	}

	for i, m := range balance {
		b.balance[i+1] = m
	}

	return b
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	b1, ok := this.balance[account1]
	if !ok || b1 < money {
		return false
	}

	_, ok = this.balance[account2]
	if !ok {
		return false
	}

	this.balance[account1] -= money
	this.balance[account2] += money

	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if b, ok := this.balance[account]; ok {
		this.balance[account] = b + money
		return true
	}

	return false
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if b, ok := this.balance[account]; ok && b >= money {
		this.balance[account] = b - money
		return true
	}

	return false
}
