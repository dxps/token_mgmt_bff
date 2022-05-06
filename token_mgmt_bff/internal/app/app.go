package app

import (
	"context"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"

	"dxps.io/token_mgmt_bff/internal/domain/logic"
	"dxps.io/token_mgmt_bff/internal/httpapi"
	"dxps.io/token_mgmt_bff/internal/infra/repo"
)

type App struct {
	httpAPI  *httpapi.API
	authnMgr *logic.AuthnMgr
	wg       *sync.WaitGroup
}

func NewApp(httpPort int) (*App, error) {

	// "inline" init for demo purposes
	clientsRepo := repo.NewClientsRepo()
	tokenFactory := logic.NewTokenFactory(10)
	authnMgr := logic.NewAuthnMgr(clientsRepo, tokenFactory)
	accountsMgr := logic.NewAccountsMgr()
	httpAPI := httpapi.NewAPI(httpPort, authnMgr, accountsMgr)
	a := App{
		httpAPI:  httpAPI,
		authnMgr: authnMgr,
		wg:       nil,
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
	// a.authnMgr.StartCleanupJob()
	return nil
}

func (a *App) Stop(stopCtx context.Context) {

	if err := a.httpAPI.Shutdown(stopCtx); err != nil {
		log.Errorf("API shutdown error: %v", err)
	} else {
		log.Info("API shutdown complete")
	}
	a.authnMgr.StopCleanupJob()
	a.wg.Done()
}
