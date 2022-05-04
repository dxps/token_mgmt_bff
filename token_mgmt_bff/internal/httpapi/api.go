package httpapi

import (
	"context"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"dxps.io/token_mgmt_bff/internal/domain/logic"
)

type API struct {
	httpServer  *http.Server
	accountsMgr *logic.AccountsMgr
}

func NewAPI(httpPort int, transferMgr *logic.AccountsMgr) *API {

	a := API{
		httpServer:  nil, // Inited below.
		accountsMgr: transferMgr,
	}
	s := http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: a.routes(),
	}
	a.httpServer = &s
	return &a
}

func logError(r *http.Request, err error) {
	log.Errorf("Failed processing HTTP request '%s %s': %s", r.Method, r.URL.String(), err)
}

func (a *API) Serve() error {

	log.Printf("HTTP API listening on port %s", a.httpServer.Addr)
	return a.httpServer.ListenAndServe()
}

func (a *API) Shutdown(stopCtx context.Context) error {

	return a.httpServer.Shutdown(stopCtx)
}
