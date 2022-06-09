package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)
		got := wallet.Balance()

		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)
		got := wallet.Balance()

		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})
}
