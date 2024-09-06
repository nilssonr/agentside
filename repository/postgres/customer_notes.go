package postgres

import (
	"context"

	"github.com/nilssonr/agentside/customer"
)

type CustomerNoteRepository struct {
	db *Queries
}

func NewCustomerNoteRepository(db *Queries) customer.NoteRepository {
	return &CustomerNoteRepository{
		db: db,
	}
}

// InsertNote implements customer.NoteRepository.
func (c *CustomerNoteRepository) InsertNote(ctx context.Context, request *customer.Note) (*customer.Note, error) {
	panic("unimplemented")
}

// GetNotes implements customer.NoteRepository.
func (c *CustomerNoteRepository) GetNotes(ctx context.Context, customerID string) ([]*customer.Note, error) {
	panic("unimplemented")
}

// GetNote implements customer.NoteRepository.
func (c *CustomerNoteRepository) GetNote(ctx context.Context, customerID string, noteID string) (*customer.Note, error) {
	panic("unimplemented")
}

// UpdateNote implements customer.NoteRepository.
func (c *CustomerNoteRepository) UpdateNote(ctx context.Context, request *customer.Note) (*customer.Note, error) {
	panic("unimplemented")
}

// DeleteNote implements customer.NoteRepository.
func (c *CustomerNoteRepository) DeleteNote(ctx context.Context, customerID string, noteID string) error {
	panic("unimplemented")
}
