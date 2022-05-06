package httpapi

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

const (
	ACCOUNT_ID = "_accountID"
)

func (a *API) routes() http.Handler {

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(respondNotFound)

	router.HandlerFunc(http.MethodPost, "/authn", a.authnHandler)

	router.Handle(http.MethodGet, "/accounts",
		authzMiddleware(a.getAccountsHandler, a.authnMgr))

	// router.HandlerFunc(http.MethodPost, "/transfer", a.transferHandler)

	// var freezePath = fmt.Sprintf("/accounts/:%s/freeze", ACCOUNT_ID)
	// router.Handle(http.MethodPost, freezePath, a.freezeAccountHandler)

	// var unfreezePath = fmt.Sprintf("/accounts/:%s/unfreeze", ACCOUNT_ID)
	// router.Handle(http.MethodPost, unfreezePath, a.unfreezeAccountHandler)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:9090"},
	}).Handler(router)

	return handler
}
