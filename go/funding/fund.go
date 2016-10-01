package funding

type Fund struct {
	balance int
}

func NewFund(initBalance int) *Fund {
	return &Fund{initBalance}
}

func (f *Fund) Balance() int {
	return f.balance
}

func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}
