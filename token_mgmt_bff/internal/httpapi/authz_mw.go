package httpapi

import (
	"net/http"
	"strings"

	"dxps.io/token_mgmt_bff/internal/errs"
	"github.com/julienschmidt/httprouter"
)

func authzMiddleware(next httprouter.Handle) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ah := r.Header.Get(AUTHZ_HEADER)
		if ah == "" || !strings.HasPrefix(ah, "Bearer ") {
			respondUnauthorized(w)
			return
		}
		token := ah[7:]

		if res := validateToken([]byte(token)); res != nil {
			switch res {
			case errs.ErrTokenInvalid:
				respondUnauthorized(w)
			case errs.ErrTokenExpired:
				respondUnauthorizedWithError(w, "token expired")
			}
			return
		}

		// Call the registered handler.
		next(w, r, ps)
	}
}
