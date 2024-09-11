package auth

import "context"

type Service interface {
	Token(ctx context.Context, request *TokenRequest) (*Token, error)
}

type GrantType string

const (
	GrantAuthorizationCode GrantType = "authorization_code"
	GrantPassword          GrantType = "password"
	GrantClientCredentials GrantType = "client_credentials"
	GrantRefreshToken      GrantType = "refresh_token"
)

type TokenRequest struct {
	GrantType    GrantType `json:"grant_type"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
}
