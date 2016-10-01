package funding

import "fmt"

type FundServer struct {
	commands chan interface{}
	fund     *Fund
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		commands: make(chan interface{}),
		fund:     NewFund(initialBalance),
	}

	go server.loop()
	return server
}

func (s *FundServer) loop() {
	for command := range s.commands {
		switch command.(type) {
		case WithdrawCommand:
			withdrawal := command.(WithdrawCommand)
			s.fund.Withdraw(withdrawal.Amount)
		case BalanceCommand:
			balance := command.(BalanceCommand)
			balance.Response <- s.fund.Balance()
		default:
			panic(fmt.Sprintf("Unknown Command: %v", command))
		}
	}
}

type WithdrawCommand struct {
	Amount int
}

type BalanceCommand struct {
	Response chan int
}

func (s *FundServer) Balance() int {
	responseChan := make(chan int)
	s.commands <- BalanceCommand{Response: responseChan}
	return <-responseChan
}

func (s *FundServer) Withdraw(amount int) {
	s.commands <- WithdrawCommand{Amount: amount}
}
