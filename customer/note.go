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
func (s *noteService) CreateNote(ctx context.Context, request *Note) (*Note, error) {
	result, err := s.noteRepository.InsertNote(ctx, request)
	if err != nil {
		s.logger.Error("failed to create customer note",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetNotes implements NoteService.
func (s *noteService) GetNotes(ctx context.Context, customerID string) ([]*Note, error) {
	result, err := s.noteRepository.GetNotes(ctx, customerID)
	if err != nil {
		s.logger.Error("failed to get customer notes",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// GetNote implements NoteService.
func (s *noteService) GetNote(ctx context.Context, customerID string, noteID string) (*Note, error) {
	result, err := s.noteRepository.GetNote(ctx, customerID, noteID)
	if err != nil {
		s.logger.Error("failed to get customer note",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// UpdateNote implements NoteService.
func (s *noteService) UpdateNote(ctx context.Context, request *Note) (*Note, error) {
	result, err := s.noteRepository.UpdateNote(ctx, request)
	if err != nil {
		s.logger.Error("failed to update customer note",
			zap.Error(err))
		return nil, err
	}

	return result, nil
}

// DeleteNote implements NoteService.
func (s *noteService) DeleteNote(ctx context.Context, customerID string, noteID string) error {
	if err := s.noteRepository.DeleteNote(ctx, customerID, noteID); err != nil {
		s.logger.Error("failed to delete customer note",
			zap.Error(err))
		return err
	}

	return nil
}
