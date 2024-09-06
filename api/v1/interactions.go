package v1

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// CreateInteraction implements ServerInterface.
func (ah AgentsideHandler) CreateInteraction(w http.ResponseWriter, r *http.Request) {
	var body CreateInteractionRequest
	if err := render.DecodeJSON(r.Body, &body); err != nil {
		handleError(w, r, err)
		return
	}
	defer r.Body.Close()

}

// GetInteraction implements ServerInterface.
func (ah AgentsideHandler) GetInteraction(w http.ResponseWriter, r *http.Request, interactionId uuid.UUID) {
	panic("unimplemented")
}

// GetInteractions implements ServerInterface.
func (ah AgentsideHandler) GetInteractions(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// CreateInteractionNote implements ServerInterface.
func (ah AgentsideHandler) CreateInteractionNote(w http.ResponseWriter, r *http.Request, interactionId uuid.UUID) {
	panic("unimplemented")
}

// GetInteractionNote implements ServerInterface.
func (ah AgentsideHandler) GetInteractionNote(w http.ResponseWriter, r *http.Request, interactionId uuid.UUID, noteId uuid.UUID) {
	panic("unimplemented")
}

// GetInteractionNotes implements ServerInterface.
func (ah AgentsideHandler) GetInteractionNotes(w http.ResponseWriter, r *http.Request, interactionId uuid.UUID) {
	panic("unimplemented")
}

// UpdateInteractionNote implements ServerInterface.
func (ah AgentsideHandler) UpdateInteractionNote(w http.ResponseWriter, r *http.Request, interactionId uuid.UUID, noteId uuid.UUID) {
	panic("unimplemented")
}

// DeleteInteractionNote implements ServerInterface.
func (ah AgentsideHandler) DeleteInteractionNote(w http.ResponseWriter, r *http.Request, interactionId uuid.UUID, noteId uuid.UUID) {
	panic("unimplemented")
}
