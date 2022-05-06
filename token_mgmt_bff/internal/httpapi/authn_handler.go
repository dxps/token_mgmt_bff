package httpapi

import (
	"encoding/gob"
	"net/http"

	"dxps.io/token_mgmt_bff/internal/domain/model"
	"dxps.io/token_mgmt_bff/internal/errs"
)

func init() {
	gob.Register(model.Client{})
}

func (a *API) authnHandler(w http.ResponseWriter, r *http.Request) {

	var c model.Client
	if err := readJSON(w, r, &c); err != nil {
		respondBadRequest(w, r, err)
	}
	t, err := a.authnMgr.Authenticate(&c)
	if err != nil {
		respondForbidden(w, r, errs.ErrInvalidCredentials)
		return
	}
	respondOKwithJSON(w, r, envelope{
		"token":      t.Value,
		"expires_at": t.ExpiresAt,
	})
}
