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
		cleanupTicker: time.NewTicker(10 * time.Second),
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
	log.Debugf("clientID '%s' got token '%v'.", client.ID, tkn)
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
	for _, t := range clientTokens {
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

func (am *AuthnMgr) RenewToken(currToken string) (*model.Token, error) {

	tknParts := strings.Split(currToken, ":")
	if len(tknParts) != 2 {
		log.Debugf("Invalid value of token '%s'.", currToken)
		return nil, errs.ErrTokenInvalid
	}
	clientID := tknParts[0]
	tkn := am.tokenFactory.NewToken(clientID)
	am.clientsTokens[clientID] = append(am.clientsTokens[clientID], tkn)
	return tkn, nil
}

func (am *AuthnMgr) StartCleanupJob() {

	go func() {
		log.Debug("AuthnMgr CleanupJob started.")
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
		log.Debug("AuthnMgr Cleanup job stopped.")
	}()
}

func (am *AuthnMgr) StopCleanupJob() {
	am.cleanupTicker.Stop()
}
