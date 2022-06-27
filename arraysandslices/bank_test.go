package arraysandslices_test

import (
	"hello/arraysandslices"
	"testing"
)

func TestBadBank(t *testing.T) {
	var (
		riya  = arraysandslices.Account{Name: "Riya", Balance: 100}
		chris = arraysandslices.Account{Name: "Chris", Balance: 75}
		adil  = arraysandslices.Account{Name: "Adil", Balance: 200}

		transactions = []arraysandslices.Transaction{
			arraysandslices.NewTransaction(chris, riya, 100),
			arraysandslices.NewTransaction(adil, chris, 25),
		}
	)

	balanceFor := func(account arraysandslices.Account) float64 {
		return arraysandslices.NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, 200, balanceFor(riya))
	AssertEqual(t, 0, balanceFor(chris))
	AssertEqual(t, 175, balanceFor(adil))
}
