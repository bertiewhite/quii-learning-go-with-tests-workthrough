package wallets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	balance := wallet.Balance()
	expected := "10 BTC"

	fmt.Printf("%d \n", balance)

	assert.Equal(t, expected, balance.String())
}

func TestWithdraw(t *testing.T) {
	wallet := Wallet{
		balance: Bitcoin(10),
	}

	err := wallet.Withdraw(Bitcoin(10))
	assert.Nil(t, err)

	balance := wallet.Balance()

	expected := "0 BTC"

	assert.Equal(t, expected, balance.String())
}

func TestErrorFromWithdraw(t *testing.T) {
	wallet := Wallet{
		balance: Bitcoin(10),
	}

	err := wallet.Withdraw(Bitcoin(20))

	assert.Equal(t, insufficientFundsError, err)
}
