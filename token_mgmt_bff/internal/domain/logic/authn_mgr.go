package logic

import (
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"dxps.io/token_mgmt_bff/internal/domain/model"
	"dxps.io/token_mgmt_bff/internal/errs"
	"dxps.io/token_mgmt_bff/internal/infra/repo"
)

type AuthnMgr struct {
	clientsRepo   *repo.ClientsRepo         // Repository of clients (and their credentials).
	tokenFactory  *TokenFactory             // Factory for generating tokens.
	clientsTokens map[string][]*model.Token // Storing per clientID the associated tokens.
	cleanupTicker *time.Ticker
}

func NewAuthnMgr(clientsRepo *repo.ClientsRepo, tokenFactory *TokenFactory) *AuthnMgr {

	return &AuthnMgr{
		clientsRepo:   clientsRepo,
		tokenFactory:  tokenFactory,
		clientsTokens: map[string][]*model.Token{},
		cleanupTicker: time.NewTicker(5 * time.Second),
	}
}

func (am *AuthnMgr) Authenticate(client *model.Client) (*model.Token, error) {

	c, err := am.clientsRepo.Get(client.ID)
	if err != nil {
		return nil, errs.ErrNotFound
	}
	if c.Secret != client.Secret {
		return nil, errs.ErrInvalidCredentials
	}
	tkn := am.tokenFactory.NewToken(client.ID)
	am.clientsTokens[client.ID] = append(am.clientsTokens[client.ID], tkn)
	log.Debugf("clientID '%s' got new token '%v'.", client.ID, tkn)
	return tkn, nil
}

func (am *AuthnMgr) ValidateToken(tokenValue string) error {

	tknParts := strings.Split(tokenValue, ":")
	if len(tknParts) != 2 {
		log.Debugf("Invalid value of token '%s'.", tokenValue)
		return errs.ErrTokenInvalid
	}
	clientTokens, ok := am.clientsTokens[tknParts[0]]
	if !ok {
		log.Debugf("There are no tokens for clientID '%s'.", tknParts[0])
		return errs.ErrTokenInvalid
	}
	log.Debugf("Validating token '%v' against '%v' ...", tokenValue, tknParts[1])
	for _, t := range clientTokens {
		log.Debugf("Token '%v' is valid? %v", t, t.IsValid())
		if t.Value == tokenValue {
			if t.IsValid() {
				return nil
			} else {
				return errs.ErrTokenExpired
			}
		}
	}
	return errs.ErrTokenInvalid
}

func (am *AuthnMgr) StartCleanupJob() {

	go func() {
		for range am.cleanupTicker.C {
			for c, ts := range am.clientsTokens {
				for i, t := range ts {
					if !t.IsValid() {
						// Removing it from the slice in an efficient way
						// with the caveat of not maintaining the order.
						sz := len(ts) - 1
						ts[i] = ts[sz]
						ts = ts[:sz]
						am.clientsTokens[c] = ts
						log.Debugf("Removed an expired token for clientID '%s', now tokens %v", c, ts)
					}
				}
			}
		}
	}()
}

func (am *AuthnMgr) StopCleanupJob() {
	am.cleanupTicker.Stop()
}
