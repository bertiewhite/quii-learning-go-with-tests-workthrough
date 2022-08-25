package wallets

import "fmt"

type Bitcoin int

var insufficientFundsError = fmt.Errorf("insufficient funds in the account")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)

}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return insufficientFundsError
	}
	w.balance -= amount
	return nil
}
