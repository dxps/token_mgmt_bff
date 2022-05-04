package httpapi

import (
	"net/http"
)

func respondOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func respondOKwithJSON(w http.ResponseWriter, r *http.Request, data any) {

	if err := writeJSON(w, http.StatusOK, data, nil); err != nil {
		respondInternalServerError(w, r, err)
	}
}

func respondError(w http.ResponseWriter, r *http.Request, status int, message interface{}) {

	env := envelope{"error": message}
	if err := writeJSON(w, status, env, nil); err != nil {
		logError(r, err)
		w.WriteHeader(500)
	}
}

func respondInternalServerError(w http.ResponseWriter, r *http.Request, err error) {

	logError(r, err)
	message := "the server encountered a problem and could not process your request"
	respondError(w, r, http.StatusInternalServerError, message)
}

func respondBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	respondError(w, r, http.StatusBadRequest, err.Error())
}

func respondNotFound(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	respondError(w, r, http.StatusNotFound, message)
}

func respondUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
}

func respondUnauthorizedWithError(w http.ResponseWriter, message interface{}) {
	e := envelope{"error": message}
	if err := writeJSON(w, http.StatusUnauthorized, e, nil); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func respondForbidden(w http.ResponseWriter, r *http.Request, err error) {
	respondError(w, r, http.StatusForbidden, err.Error())
}
