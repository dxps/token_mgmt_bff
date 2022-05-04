package model

import (
	"dxps.io/token_mgmt_bff/internal/domain/model/ccy"
)

type Account struct {
	ID       string       `json:"id"`        // The external ID of the account.
	Name     string       `json:"name"`      // The name of the account. It plays also the label role in the UI.
	Balance  int64        `json:"balance"`   // The amount of money that exist in the account.
	Currency ccy.Currency `json:"currency"`  // The currency of the account.
	IsFrozen bool         `json:"is_frozen"` // A frozen account cannot be used for debit transactions from it.
}

func NewAccount(id, name string, balance int64, currency ccy.Currency) *Account {

	a := Account{
		ID:       id,
		Name:     name,
		Balance:  balance,
		Currency: currency,
		IsFrozen: false,
	}
	return &a
}
