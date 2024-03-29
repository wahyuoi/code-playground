package funding

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkFund(b *testing.B) {
	fund := NewFund(b.N)

	for i := 0; i < b.N; i++ {
		fund.Withdraw(1)
	}

	if fund.Balance() != 0 {
		b.Error("Balance was not zero:", fund.Balance())
	}
}

const WORKERS = 10

func BenchmarkWithdraw(b *testing.B) {
	if b.N < WORKERS {
		return
	}

	server := NewFundServer(b.N)

	dollarsPerFounder := b.N / WORKERS

	fmt.Println("WORKERS:", WORKERS)
	fmt.Println("b.N:", b.N)
	fmt.Println("dollarsPerFounder:", dollarsPerFounder)

	var wg sync.WaitGroup

	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pizzaTime := false
			for ii := 0; ii < dollarsPerFounder; ii++ {
				server.Transact(func(managedValue interface{}) {
					fund := managedValue.(*Fund)
					if fund.Balance() <= 10 {
						pizzaTime = true
						return
					}
					fund.Withdraw(1)
				})
				if pizzaTime {
					break
				}
			}
		}()
	}

	wg.Wait()

	balance := server.Balance()

	if balance != 10 {
		b.Error("Balance was not ten:", balance)
	}
}
