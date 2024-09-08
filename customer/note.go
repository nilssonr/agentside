package customer

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type Note struct {
	ID             string    `json:"id"`
	Note           string    `json:"note"`
	CustomerID     string    `json:"customerId"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	LastModifiedBy string    `json:"lastModifiedBy"`
}

type NoteService interface {
	CreateNote(ctx context.Context, request *Note) (*Note, error)
	GetNotes(ctx context.Context, customerID string) ([]*Note, error)
	GetNote(ctx context.Context, customerID, noteID string) (*Note, error)
	UpdateNote(ctx context.Context, request *Note) (*Note, error)
	DeleteNote(ctx context.Context, customerID, noteID string) error
}

type noteService struct {
	noteRepository NoteRepository
	logger         *zap.Logger
}

func NewNoteService(r NoteRepository, l *zap.Logger) NoteService {
	return &noteService{
		noteRepository: r,
		logger:         l,
	}
}

// CreateNote implements NoteService.
func (n *noteService) CreateNote(ctx context.Context, request *Note) (*Note, error) {
	result, err := n.noteRepository.InsertNote(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetNotes implements NoteService.
func (n *noteService) GetNotes(ctx context.Context, customerID string) ([]*Note, error) {
	result, err := n.noteRepository.GetNotes(ctx, customerID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetNote implements NoteService.
func (n *noteService) GetNote(ctx context.Context, customerID string, noteID string) (*Note, error) {
	result, err := n.noteRepository.GetNote(ctx, customerID, noteID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateNote implements NoteService.
func (n *noteService) UpdateNote(ctx context.Context, request *Note) (*Note, error) {
	result, err := n.noteRepository.UpdateNote(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteNote implements NoteService.
func (n *noteService) DeleteNote(ctx context.Context, customerID string, noteID string) error {
	if err := n.noteRepository.DeleteNote(ctx, customerID, noteID); err != nil {
		return err
	}

	return nil
}
