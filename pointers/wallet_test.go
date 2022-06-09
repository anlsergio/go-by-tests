package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	want := 10
	got := wallet.Balance()

	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}
