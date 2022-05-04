package httpapi

import (
	"net/http"

	"dxps.io/token_mgmt_bff/internal/infra/repo"
	"github.com/julienschmidt/httprouter"
)

func (a *API) freezeAccountHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	accountID := ps.ByName(ACCOUNT_ID)
	if err := a.accountsMgr.Freeze(accountID); err != nil {
		switch err {
		case repo.ErrAccountNotFound:
			respondBadRequest(w, r, err)
		default:
			respondInternalServerError(w, r, err)
		}
		return
	}
	respondOK(w)
}

func (a *API) unfreezeAccountHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	accountID := ps.ByName(ACCOUNT_ID)
	if err := a.accountsMgr.Unfreeze(accountID); err != nil {
		switch err {
		case repo.ErrAccountNotFound:
			respondBadRequest(w, r, err)
		default:
			respondInternalServerError(w, r, err)
		}
		return
	}
	respondOK(w)
}
