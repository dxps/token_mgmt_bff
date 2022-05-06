package httpapi

import (
	"net/http"
	"strings"

	"dxps.io/token_mgmt_bff/internal/domain/logic"
	"dxps.io/token_mgmt_bff/internal/errs"
	"github.com/julienschmidt/httprouter"
)

func authzMiddleware(next httprouter.Handle, authnMgr *logic.AuthnMgr) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ah := r.Header.Get(AUTHZ_HEADER)
		if ah == "" || !strings.HasPrefix(ah, "Bearer ") {
			respondUnauthorized(w)
			return
		}
		token := ah[7:]

		if err := authnMgr.ValidateToken(token); err != nil {
			switch err {
			case errs.ErrTokenInvalid:
				respondForbidden(w, r, errs.ErrTokenInvalid)
				return
			case errs.ErrTokenExpired:
				respondForbidden(w, r, errs.ErrTokenExpired)
				return
			}

		}

		// Call the registered handler.
		next(w, r, ps)
	}
}
