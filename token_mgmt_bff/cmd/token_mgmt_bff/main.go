package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"dxps.io/token_mgmt_bff/internal/app"
)

const (
	HTTP_PORT = 9093
)

func main() {

	initLogging()

	wg := &sync.WaitGroup{}

	app, err := app.NewApp(HTTP_PORT)
	if err != nil {
		log.Fatalf("App init failed: %s", err)
	}

	if err := app.Start(wg); err != nil {
		log.Fatalf("Startup failed: %v", err)
	}

	// Graceful shutdown setup & usage.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)

	// Waiting for shutdown signal.
	<-signalChan
	log.Info("Shutting down ...")

	stopCtx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	app.Stop(stopCtx)

	wg.Wait()
	log.Println("Exit now")
}

func initLogging() {
	// For now, we just need the text format. Later, we shall use JSON.
	// log.SetFormatter(&log.JSONFormatter{DisableHTMLEscape: true})
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, PadLevelText: true})
	log.SetLevel(log.DebugLevel)
}
