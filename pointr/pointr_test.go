package pointr

import (
	"testing"
)

//https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError:= func (t testing.TB,err error, want string)  {
		t.Helper()
		if err == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		if err.Error() != want {
			t.Errorf("got %q want %q ", err, want)
		}
	}
	assertNoError := func (t testing.TB, err error)  {
		t.Helper()
		if err != nil {
			t.Fatal("got an error but didn't want one")
		}
	}
	t.Run("test deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("test withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(50))
		assertBalance(t, wallet, Bitcoin(50))
		assertNoError(t,err)

	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t,err, ErrInsufficientFunds.Error())

		// if err == nil {
		// 	t.Error("wanted an error but didn't get one")
		// }
	})

}
