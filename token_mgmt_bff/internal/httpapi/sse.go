package httpapi

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func prepareHeadersForSSE(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func respondInitial(w http.ResponseWriter) (int, error) {
	return fmt.Fprintf(w, "\r")
}

func (a *API) tokenStreamHandler(w http.ResponseWriter, r *http.Request) {

	prepareHeadersForSSE(w)
	// http.Flusher allows this handler to flush buffered data
	// to ResponseWriter, until the client closes the connection.
	flusher, ok := w.(http.Flusher)
	if !ok {
		log.Errorf("[tokenStreamHandler] ResponseWriter cannot be flushed.")
		return
	}

	// Initial response right away, so that Browser sees the established connection
	// when the UI loads and establish the SSE connection with this service.
	if _, err := respondInitial(w); err != nil {
		log.Errorf("Failed to send initial response: %v\n", err)
	}
	flusher.Flush()

	initTkn := r.URL.Query().Get("token")
	log.Debugf("[tokenStreamHandler] Initial token is '%v'.", initTkn)
	for {
		tkn, err := a.authnMgr.RenewToken(initTkn)
		if err != nil {
			log.Errorf("[tokenStreamHandler] Renew token failed: %v", err)
			fmt.Fprintf(w, "error: %s\n\n", err)
		} else {
			log.Debugf("[tokenStreamHandler] Renewed token: %v", tkn.Value)
			fmt.Fprintf(w, "data: %s\n\n", tkn.Value)
		}
		flusher.Flush()

		time.Sleep(5 * time.Second)
	}
}
