package httpapi

import (
	"net/http"

	corsHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	// "github.com/rs/cors"
)

const (
	ACCOUNT_ID = "_accountID"
)

var (
	CORS_ALLOWED_ORIGINS []string = []string{"http://localhost:9090"}
	CORS_ALLOWED_HEADERS []string = []string{"Origin", "Authorization", "Content-Type"}
	CORS_MAX_AGE         int      = 3600
)

func (a *API) routes() http.Handler {

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(respondNotFound)

	r.HandleFunc("/authn", a.authnHandler)

	r.HandleFunc("/accounts", a.getAccountsHandler)

	r.HandleFunc("/sse/stream", a.tokenStreamHandler)

	// handler := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
	// 	AllowedHeaders: []string{"Origin", "Authorization", "Content-Type"},
	// }).Handler(r)
	// return handler

	corsMaxAge := corsHandlers.MaxAge(CORS_MAX_AGE)
	corsOrigins := corsHandlers.AllowedOrigins(CORS_ALLOWED_ORIGINS)
	corsHeaders := corsHandlers.AllowedHeaders(CORS_ALLOWED_HEADERS)
	corsHandler := corsHandlers.CORS(corsMaxAge, corsOrigins, corsHeaders)
	return corsHandler(r)
}
