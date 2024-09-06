package v1

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nilssonr/agentside/interaction"
)

type InteractionHandler struct {
	InteractionService interaction.Service
	// InteractionnoteService interaction.NoteService
}

func NewInteractionHandler(
	interactionService interaction.Service,
) InteractionHandler {
	return InteractionHandler{
		InteractionService: interactionService,
	}
}

// CreateInteraction implements ServerInterface.
func (h InteractionHandler) CreateInteraction(w http.ResponseWriter, r *http.Request) {
	var body CreateInteractionRequest
	if err := render.DecodeJSON(r.Body, &body); err != nil {
		handleError(w, r, err)
		return
	}
	defer r.Body.Close()

}

// GetInteraction implements ServerInterface.
func (h InteractionHandler) GetInteraction(w http.ResponseWriter, r *http.Request, interactionId uuid.UUID) {
	panic("unimplemented")
}

// GetInteractions implements ServerInterface.
func (h InteractionHandler) GetInteractions(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// CreateInteractionNote implements ServerInterface.
func (h InteractionHandler) CreateInteractionNote(w http.ResponseWriter, r *http.Request, interactionId uuid.UUID) {
	panic("unimplemented")
}

// GetInteractionNote implements ServerInterface.
func (h InteractionHandler) GetInteractionNote(w http.ResponseWriter, r *http.Request, interactionId uuid.UUID, noteId uuid.UUID) {
	panic("unimplemented")
}

// GetInteractionNotes implements ServerInterface.
func (h InteractionHandler) GetInteractionNotes(w http.ResponseWriter, r *http.Request, interactionId uuid.UUID) {
	panic("unimplemented")
}

// UpdateInteractionNote implements ServerInterface.
func (h InteractionHandler) UpdateInteractionNote(w http.ResponseWriter, r *http.Request, interactionId uuid.UUID, noteId uuid.UUID) {
	panic("unimplemented")
}

// DeleteInteractionNote implements ServerInterface.
func (h InteractionHandler) DeleteInteractionNote(w http.ResponseWriter, r *http.Request, interactionId uuid.UUID, noteId uuid.UUID) {
	panic("unimplemented")
}
