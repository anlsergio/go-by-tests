package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	}

	assertError := func(t testing.TB, want string, got error) {
		t.Helper()

		if got == nil {
			t.Fatal("expected error but didn't get any")
		}

		if want != got.Error() {
			t.Errorf("want %q, got %q", want, got)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(startingBalance + 1))

		wantErr := "insufficient funds"
		assertError(t, wantErr, err)
		assertBalance(t, wallet, startingBalance)
	})
}
