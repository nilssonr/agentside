package postgres

import (
	"context"
	"fmt"

	"github.com/nilssonr/agentside/customer"
	"github.com/nilssonr/agentside/repository/postgres/sqlc"
)

type CustomerNoteRepository struct {
	DB *sqlc.Queries
}

func NewCustomerNoteRepository(db *sqlc.Queries) customer.NoteRepository {
	return &CustomerNoteRepository{
		DB: db,
	}
}

// InsertNote implements customer.NoteRepository.
func (r *CustomerNoteRepository) InsertNote(ctx context.Context, request *customer.Note) (*customer.Note, error) {
	arg := sqlc.InsertCustomerNoteParams{
		Note:           request.Note,
		CustomerID:     request.CustomerID,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	row, err := r.DB.InsertCustomerNote(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &customer.Note{
		ID:             row.ID,
		Note:           row.Note,
		CustomerID:     row.CustomerID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// GetNotes implements customer.NoteRepository.
func (r *CustomerNoteRepository) GetNotes(ctx context.Context, customerID string) ([]*customer.Note, error) {
	rows, err := r.DB.GetCustomerNotes(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := make([]*customer.Note, 0, len(rows))
	for _, v := range rows {
		result = append(result, &customer.Note{
			ID:             v.ID,
			Note:           v.Note,
			CustomerID:     v.CustomerID,
			LastModifiedAt: v.LastModifiedAt.Time,
			LastModifiedBy: v.LastModifiedBy,
		})
	}

	return result, nil
}

// GetNote implements customer.NoteRepository.
func (r *CustomerNoteRepository) GetNote(ctx context.Context, customerID, noteID string) (*customer.Note, error) {
	arg := sqlc.GetCustomerNoteParams{
		CustomerID: customerID,
		ID:         noteID,
	}

	row, err := r.DB.GetCustomerNote(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &customer.Note{
		ID:             row.ID,
		Note:           row.Note,
		CustomerID:     row.CustomerID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// UpdateNote implements customer.NoteRepository.
func (r *CustomerNoteRepository) UpdateNote(ctx context.Context, request *customer.Note) (*customer.Note, error) {
	arg := sqlc.UpdateCustomerNoteParams{
		CustomerID:     request.CustomerID,
		ID:             request.ID,
		Note:           request.Note,
		LastModifiedAt: mustCreateTime(request.LastModifiedAt),
		LastModifiedBy: request.LastModifiedBy,
	}

	row, err := r.DB.UpdateCustomerNote(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}

	result := &customer.Note{
		ID:             row.ID,
		Note:           row.Note,
		CustomerID:     row.CustomerID,
		LastModifiedAt: row.LastModifiedAt.Time,
		LastModifiedBy: row.LastModifiedBy,
	}

	return result, nil
}

// DeleteNote implements customer.NoteRepository.
func (r *CustomerNoteRepository) DeleteNote(ctx context.Context, customerID, noteID string) error {
	arg := sqlc.DeleteCustomerNoteParams{
		CustomerID: customerID,
		ID:         noteID,
	}

	if err := r.DB.DeleteCustomerNote(ctx, arg); err != nil {
		return fmt.Errorf("repository: %w", err)
	}

	return nil
}
