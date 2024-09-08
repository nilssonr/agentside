package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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
	var body createQueueRequest
	if err := render.Decode(r, &body); err != nil {
		handleError(w, r, err)
		return
	}

	request := &queue.Queue{
		Name:           body.Name,
		TenantID:       tenantID(r),
		LastModifiedAt: time.Now(),
		LastModifiedBy: userID(r),
	}

	result, err := h.QueueService.CreateQueue(r.Context(), request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h queueHandler) getQueues(w http.ResponseWriter, r *http.Request) {
	result, err := h.QueueService.GetQueues(r.Context(), tenantID(r))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h queueHandler) getQueue(w http.ResponseWriter, r *http.Request) {
	result, err := h.QueueService.GetQueue(r.Context(), tenantID(r), chi.URLParam(r, "queueID"))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h queueHandler) updateQueue(w http.ResponseWriter, r *http.Request) {
	var body updateQueueRequest
	if err := render.Decode(r, &body); err != nil {
		handleError(w, r, err)
		return
	}

	request := &queue.Queue{
		ID:             chi.URLParam(r, "queueID"),
		Name:           body.Name,
		TenantID:       tenantID(r),
		LastModifiedAt: time.Now(),
		LastModifiedBy: userID(r),
	}

	result, err := h.QueueService.UpdateQueue(r.Context(), request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h queueHandler) deleteQueue(w http.ResponseWriter, r *http.Request) {
	if err := h.QueueService.DeleteQueue(r.Context(), tenantID(r), chi.URLParam(r, "queueID")); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}
