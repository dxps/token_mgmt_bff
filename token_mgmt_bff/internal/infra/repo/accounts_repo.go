package repo

import (
	"errors"

	"dxps.io/token_mgmt_bff/internal/domain/model"
	"dxps.io/token_mgmt_bff/internal/domain/model/ccy"
)

var (
	ErrAccountNotFound = errors.New("account is not found")
)

type AccountsRepo struct {
	store []*model.Account
}

func NewAccountsRepo() *AccountsRepo {

	store := make([]*model.Account, 0)
	// Just for demo purposes only.
	store = append(store, model.NewAccount("1", "Main Account", 1_000, ccy.USD))
	store = append(store, model.NewAccount("2", "Savings Account", 2_000, ccy.USD))
	r := AccountsRepo{store}
	return &r
}

func (r *AccountsRepo) GetAll() []*model.Account {

	res := make([]*model.Account, 0)
	for _, a := range r.store {
		res = append(res, a)
	}
	return res
}

func (r *AccountsRepo) Get(accountID string) (*model.Account, error) {

	for _, a := range r.store {
		if a.ID == accountID {
			return a, nil
		}
	}
	return nil, ErrAccountNotFound
}

func (r *AccountsRepo) UpdateBalance(accountID string, balance int64) error {

	a, err := r.Get(accountID)
	if err != nil {
		return err
	}
	a.Balance = balance
	return nil
}

func (r *AccountsRepo) UpdateFreeze(accountID string, freeze bool) error {

	a, err := r.Get(accountID)
	if err != nil {
		return err
	}
	a.IsFrozen = freeze
	return nil
}
