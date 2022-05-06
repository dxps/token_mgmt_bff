package logic

import (
	"fmt"
	"time"

	"dxps.io/token_mgmt_bff/internal/domain/model"
	"github.com/google/uuid"
)

type TokenFactory struct {
	tokenLifespan int // Number of seconds a token is valid.
}

// NewTokenFactory creates a TokenFactory, with tokenLifespan meaning
// the number of seconds a generated token is valid.
func NewTokenFactory(tokenLifespan int) *TokenFactory {
	return &TokenFactory{tokenLifespan}
}

func (tf *TokenFactory) NewToken(clientID string) *model.Token {
	return &model.Token{
		Value:     fmt.Sprintf("%s:%s", clientID, uuid.NewString()),
		ExpiresAt: time.Now().Unix() + int64(tf.tokenLifespan),
	}
}
