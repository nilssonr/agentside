package interaction

import "context"

type Repository interface {
	InsertInteraction(ctx context.Context, request *Interaction) (*Interaction, error)
	GetInteractions(ctx context.Context, tenantID string) ([]*Interaction, error)
	GetInteraction(ctx context.Context, tenantID, interactionID string) (*Interaction, error)
}
