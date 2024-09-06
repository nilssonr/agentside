package postgres

import (
	"context"

	"github.com/nilssonr/agentside/interaction"
)

type InteractionRepository struct {
	db *Queries
}

func NewInteractionRepository(db *Queries) interaction.Repository {
	return &InteractionRepository{
		db: db,
	}
}

// InsertInteraction implements interaction.Repository.
func (i *InteractionRepository) InsertInteraction(ctx context.Context, request *interaction.Interaction) (*interaction.Interaction, error) {
	arg := InsertInteractionParams{
		Type:            string(request.Type),
		QueueID:         request.QueueID,
		State:           string(request.State),
		StateModifiedAt: mustCreateTime(request.StateModifiedAt),
		TenantID:        request.TenantID,
		CreatedAt:       mustCreateTime(request.CreatedAt),
	}

	row, err := i.db.InsertInteraction(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result interaction.Interaction
	result.ID = row.ID
	result.Type = interaction.Type(row.Type)
	result.QueueID = row.QueueID
	result.State = interaction.State(row.State)
	result.StateModifiedAt = row.StateModifiedAt.Time
	result.TenantID = row.TenantID
	result.CreatedAt = row.CreatedAt.Time

	if row.UserID.Valid {
		result.UserID = row.UserID.String
	}

	return &result, nil
}

// GetInteractions implements interaction.Repository.
func (i *InteractionRepository) GetInteractions(ctx context.Context, tenantID string) ([]*interaction.Interaction, error) {
	panic("unimplemented")
}

// GetInteraction implements interaction.Repository.
func (i *InteractionRepository) GetInteraction(ctx context.Context, tenantID string, interactionID string) (*interaction.Interaction, error) {
	panic("unimplemented")
}
