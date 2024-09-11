package auth

import (
	"context"

	"github.com/nilssonr/agentside/user"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type TokenService interface {
	Token(ctx context.Context, user *user.User) (*Token, error)
}
