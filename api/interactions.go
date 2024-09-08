package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nilssonr/agentside/interaction"
)

type interactionHandler struct {
	InteractionService interaction.Service
}

func (h interactionHandler) Register(r chi.Router) {
	r.Route("/interactions", func(r chi.Router) {
		r.Post("/", h.createInteraction)
		r.Get("/", h.getInteractions)

		r.Route("/{interactionID}", func(r chi.Router) {
			r.Get("/", h.getInteraction)
		})
	})
}

type createInteractionRequest struct {
}

type updateInteractionRequest struct {
}

func (h interactionHandler) createInteraction(w http.ResponseWriter, r *http.Request) {

}

func (h interactionHandler) getInteractions(w http.ResponseWriter, r *http.Request) {

}

func (h interactionHandler) getInteraction(w http.ResponseWriter, r *http.Request) {

}
