package data

import "time"

const (
	ScopeActivation     = "activation"
	ScopeAuthentication = "authentication"
)

type Token struct {
	PlainText string    `json:"plaintext"`
	Hash      []byte    `json:"hash"`
	UserID    int64     `json:"user_id"`
	Scope     string    `json:"scope"`
	ExpiresAt time.Time `json:"expires_at"`
}
