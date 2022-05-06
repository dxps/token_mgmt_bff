package httpapi

// "github.com/julienschmidt/httprouter"

// func authzMiddleware(next httprouter.Handle, authnMgr *logic.AuthnMgr) httprouter.Handle {
// func authzMiddleware(next httprouter.Handle, authnMgr *logic.AuthnMgr) httprouter.Handle {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		ah := r.Header.Get(AUTHZ_HEADER)
// 		if ah == "" || !strings.HasPrefix(ah, "Bearer ") {
// 			respondUnauthorized(w)
// 			return
// 		}
// 		token := ah[7:]

// 		if err := authnMgr.ValidateToken(token); err != nil {
// 			switch err {
// 			case errs.ErrTokenInvalid:
// 				respondForbidden(w, r, errs.ErrTokenInvalid)
// 				return
// 			case errs.ErrTokenExpired:
// 				respondForbidden(w, r, errs.ErrTokenExpired)
// 				return
// 			}

// 		}

// 		// Call the registered handler.
// 		next(w, r, ps)
// 	}
// }
