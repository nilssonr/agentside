package postgres

import (
	"context"
	"fmt"

	"github.com/nilssonr/agentside/interaction"
	"github.com/nilssonr/agentside/repository/postgres/sqlc"
)

type InteractionRepository struct {
	DB *sqlc.Queries
}

func NewInteractionRepository(db *sqlc.Queries) interaction.Repository {
	return &InteractionRepository{
		DB: db,
	}
}

// InsertInteraction implements interaction.Repository.
func (r *InteractionRepository) InsertInteraction(ctx context.Context, request *interaction.Interaction) (*interaction.Interaction, error) {
	arg := sqlc.InsertInteractionParams{
		Type:            string(request.Type),
		QueueID:         request.QueueID,
		State:           string(request.State),
		StateModifiedAt: mustCreateTime(request.StateModifiedAt),
		TenantID:        request.TenantID,
		CreatedAt:       mustCreateTime(request.CreatedAt),
	}

	row, err := r.DB.InsertInteraction(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &interaction.Interaction{
		ID:              row.ID,
		Type:            interaction.Type(row.Type),
		QueueID:         row.QueueID,
		State:           interaction.State(row.State),
		StateModifiedAt: row.StateModifiedAt.Time,
		UserID:          row.UserID.String,
		TenantID:        row.TenantID,
		CreatedAt:       row.CreatedAt.Time,
	}

	return result, nil
}

// GetInteractions implements interaction.Repository.
func (r *InteractionRepository) GetInteractions(ctx context.Context, tenantID string) ([]*interaction.Interaction, error) {
	rows, err := r.DB.GetInteractions(ctx, tenantID)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := make([]*interaction.Interaction, 0, len(rows))
	for _, v := range rows {
		result = append(result, &interaction.Interaction{
			ID:              v.ID,
			Type:            interaction.Type(v.Type),
			QueueID:         v.QueueID,
			State:           interaction.State(v.State),
			StateModifiedAt: v.StateModifiedAt.Time,
			UserID:          v.UserID.String,
			TenantID:        v.TenantID,
			CreatedAt:       v.CreatedAt.Time,
		})
	}

	return result, nil
}

// GetInteraction implements interaction.Repository.
func (r *InteractionRepository) GetInteraction(ctx context.Context, tenantID, interactionID string) (*interaction.Interaction, error) {
	arg := sqlc.GetInteractionParams{
		TenantID: tenantID,
		ID:       interactionID,
	}

	row, err := r.DB.GetInteraction(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &interaction.Interaction{
		ID:              row.ID,
		Type:            interaction.Type(row.Type),
		QueueID:         row.QueueID,
		State:           interaction.State(row.State),
		StateModifiedAt: row.StateModifiedAt.Time,
		UserID:          row.UserID.String,
		TenantID:        row.TenantID,
		CreatedAt:       row.CreatedAt.Time,
	}

	return result, nil
}
