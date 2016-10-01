package funding

import "testing"

func BenchmarkFund(b *testing.B) {
	fund := NewFund(b.N)

	for i := 0; i < b.N; i++ {
		fund.Withdraw(1)
	}

	if fund.Balance() != 0 {
		b.Error("Balance was not zero:", fund.Balance())
	}
}
