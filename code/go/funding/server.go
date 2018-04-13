package funding

type Transactor func(interface{})

type FundServer struct {
	commands chan TransactionCommand
	fund     *Fund
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		commands: make(chan TransactionCommand),
		fund:     NewFund(initialBalance),
	}

	go server.loop()
	return server
}

func (s *FundServer) Transact(transactor Transactor) {
	command := TransactionCommand{
		Transactor: transactor,
		Done:       make(chan bool),
	}
	s.commands <- command
	<-command.Done
}

func (s *FundServer) loop() {
	for transaction := range s.commands {
		transaction.Transactor(s.fund)
		transaction.Done <- true
	}
}

type TransactionCommand struct {
	Transactor Transactor
	Done       chan bool
}

func (s *FundServer) Balance() int {
	var balance int
	s.Transact(func(managedValue interface{}) {
		fund := managedValue.(*Fund)
		balance = fund.Balance()
	})
	return balance
}

func (s *FundServer) Withdraw(amount int) {
	s.Transact(func(managedValue interface{}) {
		fund := managedValue.(*Fund)
		fund.Withdraw(amount)
	})
}
