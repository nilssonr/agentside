package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nilssonr/agentside/queue"
)

type queueHandler struct {
	QueueService queue.Service
}

func (h queueHandler) Register(r chi.Router) {
	r.Route("/queues", func(r chi.Router) {
		r.Post("/", h.createQueue)
		r.Get("/", h.getQueues)
		r.Route("/{queueID}", func(r chi.Router) {
			r.Get("/", h.getQueue)
			r.Put("/", h.updateQueue)
			r.Delete("/", h.deleteQueue)
		})
	})
}

type createQueueRequest struct {
	Name string `json:"name"`
}

type updateQueueRequest struct {
	Name string `json:"name"`
}

func (h queueHandler) createQueue(w http.ResponseWriter, r *http.Request) {

}

func (h queueHandler) getQueues(w http.ResponseWriter, r *http.Request) {

}

func (h queueHandler) getQueue(w http.ResponseWriter, r *http.Request) {

}

func (h queueHandler) updateQueue(w http.ResponseWriter, r *http.Request) {

}

func (h queueHandler) deleteQueue(w http.ResponseWriter, r *http.Request) {

}
