package app

import (
	"context"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"

	"dxps.io/token_mgmt_bff/internal/domain/logic"
	"dxps.io/token_mgmt_bff/internal/httpapi"
)

type App struct {
	httpAPI *httpapi.API
	wg      *sync.WaitGroup
}

func NewApp(httpPort int) (*App, error) {

	// "inline" init for demo purposes
	txnMgr := logic.NewAccountsMgr()
	httpAPI := httpapi.NewAPI(httpPort, txnMgr)
	a := App{
		httpAPI: httpAPI,
		wg:      nil,
	}
	return &a, nil
}

func (a *App) Start(wg *sync.WaitGroup) error {

	a.wg = wg
	go func() {
		if err := a.httpAPI.Serve(); err != http.ErrServerClosed {
			log.Fatalf("HTTP API listening failure: %s", err)
		}
	}()
	a.wg.Add(1)
	return nil
}

func (a *App) Stop(stopCtx context.Context) {

	if err := a.httpAPI.Shutdown(stopCtx); err != nil {
		log.Errorf("API shutdown error: %v", err)
	} else {
		log.Info("API shutdown complete")
	}
	a.wg.Done()
}
