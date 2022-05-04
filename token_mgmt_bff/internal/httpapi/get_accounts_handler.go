package httpapi

import "net/http"

func (a *API) getAccountsHandler(w http.ResponseWriter, r *http.Request) {

	accounts := a.accountsMgr.GetAccounts()
	respondOKwithJSON(w, r, accounts)
}
