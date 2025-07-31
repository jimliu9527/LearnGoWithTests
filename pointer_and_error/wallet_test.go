package pointeranderror

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' but want '%s'", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but want one")
		}

		if got != want {
			t.Errorf("got '%s' but want '%s'", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		if got != nil {
			t.Errorf("got an error but didn't want one ")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(0)}
		wallet.Deposit(Bitcoin(10))
		fmt.Println("address of balance in test is", &wallet.balance)

		got := wallet.Balance()
		want := Bitcoin(10)
		assertBalance(t, got, want)
	})

	t.Run("Withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))
		got := wallet.Balance()
		want := Bitcoin(5)

		assertNoError(t, err)
		assertBalance(t, got, want)
	})

	t.Run("Withdraw insufficient founds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(100)
		assertError(t, err, InsufficientFundsError)
	})

}
