package model

import (
	"time"
)

type Token struct {
	Value     string // The value of the token.
	ExpiresAt int64  // The expiry moment, stored as unix epoch timestamp.
}

func (t *Token) IsValid() bool {
	now := time.Now().Unix()
	return now < t.ExpiresAt
}
