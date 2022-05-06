package httpapi

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *API) getAccountsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	accounts := a.accountsMgr.GetAccounts()
	respondOKwithJSON(w, r, accounts)
}
