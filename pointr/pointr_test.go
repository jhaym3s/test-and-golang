package pointr

import (
	"fmt"
	"testing"
)


func TestWallet(t *testing.T) {
	t.Run("test deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
	
		got := wallet.Balance()
	
		fmt.Printf("address of balance in test is %v \n", &wallet.balance)
	
		want := Bitcoin(10)
	
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("test withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		wallet.Withdraw(Bitcoin(50))
		got :=  wallet.Balance()
		want := Bitcoin(50)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})

	
}