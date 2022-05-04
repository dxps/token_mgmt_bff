package httpapi

import (
	"encoding/gob"
	"net/http"

	"dxps.io/token_mgmt_bff/internal/domain/logic"
)

func init() {
	gob.Register(TransferPayload{})
}

type TransferPayload struct {
	SourceAccountID string `json:"source_account_id"`
	TargetAccountID string `json:"target_account_id"`
	Amount          int64  `json:"amount"`
}

func (a *API) transferHandler(w http.ResponseWriter, r *http.Request) {

	var tp TransferPayload
	if err := readJSON(w, r, &tp); err != nil {
		respondBadRequest(w, r, err)
	}
	if err := a.accountsMgr.Transfer(tp.SourceAccountID, tp.TargetAccountID, tp.Amount); err != nil {
		switch err {
		case logic.ErrTxnAccountInsufficientFunds, logic.ErrTxnAccountIsFrozen,
			logic.ErrTxnAccountSourceNotFound, logic.ErrTxnAccountTargetNotFound:
			respondForbidden(w, r, err)
		default:
			respondInternalServerError(w, r, err)
		}
		return
	}
	respondOK(w)
}
