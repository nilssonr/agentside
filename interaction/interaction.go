package interaction

import (
	"context"
	"time"
)

type Type string

const (
	TypeVoice      Type = "voice"
	TypeEmail      Type = "email"
	TypeChat       Type = "chat"
	TypeThirdParty Type = "third_party"
)

type State string

const (
	StateQueued   State = "queued"
	StateHandling State = "handling"
	StateHeld     State = "held"
	StateEnded    State = "ended"
)

type Interaction struct {
	ID              string    `json:"id"`
	Type            Type      `json:"type"`
	QueueID         string    `json:"queueId"`
	State           State     `json:"state"`
	StateModifiedAt time.Time `json:"stateModifiedAt"`
	UserID          string    `json:"userId"`
	TenantID        string    `json:"tenantId"`
	CreatedAt       time.Time `json:"createdAt"`
}

type Service interface {
	CreateInteraction(ctx context.Context, request *Interaction) (*Interaction, error)
	GetInteractions(ctx context.Context, tenantID string) ([]*Interaction, error)
	GetInteraction(ctx context.Context, tenantID, interactionID string) (*Interaction, error)
}

type service struct {
	interactionRepository Repository
}

func NewService(r Repository) Service {
	return &service{
		interactionRepository: r,
	}
}

// CreateInteraction implements Service.
func (s *service) CreateInteraction(ctx context.Context, request *Interaction) (*Interaction, error) {
	result, err := s.interactionRepository.InsertInteraction(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetInteractions implements Service.
func (s *service) GetInteractions(ctx context.Context, tenantID string) ([]*Interaction, error) {
	result, err := s.interactionRepository.GetInteractions(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetInteraction implements Service.
func (s *service) GetInteraction(ctx context.Context, tenantID string, interactionID string) (*Interaction, error) {
	result, err := s.interactionRepository.GetInteraction(ctx, tenantID, interactionID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
