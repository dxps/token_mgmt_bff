package logic

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"dxps.io/token_mgmt_bff/internal/domain/model"
	"dxps.io/token_mgmt_bff/internal/infra/repo"
)

var (
	ErrTxnAccountSourceNotFound    = errors.New("source account is not found")
	ErrTxnAccountTargetNotFound    = errors.New("target account is not found")
	ErrTxnAccountIsFrozen          = errors.New("source account is frozen")
	ErrTxnAccountInsufficientFunds = errors.New("source account has insufficient funds")
	ErrTxnInternal                 = errors.New("internal error")
)

type AccountsMgr struct {
	accountsRepo *repo.AccountsRepo
}

func NewAccountsMgr() *AccountsMgr {

	r := repo.NewAccountsRepo()
	return &AccountsMgr{
		accountsRepo: r,
	}
}

func (am *AccountsMgr) GetAccounts() []*model.Account {
	return am.accountsRepo.GetAll()
}

func (am *AccountsMgr) Transfer(sourceAccountID, targetAccountID string, amount int64) error {

	sourceAcc, err := am.accountsRepo.Get(sourceAccountID)
	if err != nil {
		switch err {
		case repo.ErrAccountNotFound:
			return ErrTxnAccountSourceNotFound
		default:
			return ErrTxnInternal
		}
	}
	if sourceAcc.IsFrozen {
		return ErrTxnAccountIsFrozen
	}
	if sourceAcc.Balance < amount {
		return ErrTxnAccountInsufficientFunds
	}

	targetAcc, err := am.accountsRepo.Get(targetAccountID)
	if err != nil {
		switch err {
		case repo.ErrAccountNotFound:
			return ErrTxnAccountTargetNotFound
		default:
			return ErrTxnInternal
		}
	}

	// Ofc, this transfer should be an atomic change.
	if err := am.accountsRepo.UpdateBalance(sourceAcc.ID, sourceAcc.Balance-amount); err != nil {
		log.Errorf("failed to update balance on account ID '%s': %v", sourceAccountID, err)
	}
	if err := am.accountsRepo.UpdateBalance(targetAcc.ID, targetAcc.Balance+amount); err != nil {
		log.Errorf("failed to update balance on account ID '%s': %v", sourceAccountID, err)
	}

	return nil
}

func (am *AccountsMgr) Freeze(accountID string) error {
	return am.accountsRepo.UpdateFreeze(accountID, true)
}

func (am *AccountsMgr) Unfreeze(accountID string) error {
	return am.accountsRepo.UpdateFreeze(accountID, false)
}
