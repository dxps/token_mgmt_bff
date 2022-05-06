package httpapi

import (
	"net/http"
	"strings"

	"dxps.io/token_mgmt_bff/internal/errs"
)

func (a *API) getAccountsHandler(w http.ResponseWriter, r *http.Request) {

	ah := r.Header.Get(AUTHZ_HEADER)
	if ah == "" || !strings.HasPrefix(ah, "Bearer ") {
		respondUnauthorized(w)
		return
	}
	token := ah[7:]

	if err := a.authnMgr.ValidateToken(token); err != nil {
		switch err {
		case errs.ErrTokenInvalid:
			respondForbidden(w, r, errs.ErrTokenInvalid)
			return
		case errs.ErrTokenExpired:
			respondForbidden(w, r, errs.ErrTokenExpired)
			return
		}
	}

	accounts := a.accountsMgr.GetAccounts()
	respondOKwithJSON(w, r, accounts)
}
