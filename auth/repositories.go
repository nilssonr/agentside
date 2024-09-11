package auth

import "context"

type ClientRepository interface {
	InsertClient(ctx context.Context, request *Client) (*Client, error)
	GetClients(ctx context.Context, tenantID string) ([]*Client, error)
	GetClient(ctx context.Context, tenantID, clientID string) (*Client, error)
	UpdateClient(ctx context.Context, request *Client) (*Client, error)
	DeleteClient(ctx context.Context, tenantID, clientID string) error
}
