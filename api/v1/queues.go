package v1

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nilssonr/agentside/queue"
)

type QueueHandler struct {
	QueueService queue.Service
	// QueueSkillService queue.SkillService
}

func NewQueueHandler(qs queue.Service) QueueHandler {
	return QueueHandler{
		QueueService: qs,
	}
}

// CreateQueue implements ServerInterface.
func (h QueueHandler) CreateQueue(w http.ResponseWriter, r *http.Request) {
	var body CreateQueueRequest
	if err := render.DecodeJSON(r.Body, &body); err != nil {
		handleError(w, r, err)
		return
	}
	defer r.Body.Close()

	var request queue.Queue
	request.Name = body.Name
	request.TenantID = tenantID(r)
	request.LastModifiedAt = time.Now()
	request.LastModifiedBy = userID(r)

	created, err := h.QueueService.CreateQueue(r.Context(), &request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, created)
}

// GetQueues implements ServerInterface.
func (h QueueHandler) GetQueues(w http.ResponseWriter, r *http.Request) {
	result, err := h.QueueService.GetQueues(r.Context(), tenantID(r))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetQueue implements ServerInterface.
func (h QueueHandler) GetQueue(w http.ResponseWriter, r *http.Request, queueId uuid.UUID) {
	result, err := h.QueueService.GetQueue(r.Context(), tenantID(r), queueId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// UpdateQueue implements ServerInterface.
func (h QueueHandler) UpdateQueue(w http.ResponseWriter, r *http.Request, queueId uuid.UUID) {
	var body UpdateQueueRequest
	if err := render.DecodeJSON(r.Body, &body); err != nil {
		handleError(w, r, err)
		return
	}
	defer r.Body.Close()

	var request queue.Queue
	request.ID = queueId.String()
	request.Name = body.Name
	request.TenantID = tenantID(r)
	request.LastModifiedAt = time.Now()
	request.LastModifiedBy = userID(r)

	result, err := h.QueueService.UpdateQueue(r.Context(), &request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// DeleteQueue implements ServerInterface.
func (h QueueHandler) DeleteQueue(w http.ResponseWriter, r *http.Request, queueId uuid.UUID) {
	if err := h.QueueService.DeleteQueue(r.Context(), tenantID(r), queueId.String()); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}

// UpsertQueueSkill implements ServerInterface.
func (h QueueHandler) UpsertQueueSkill(w http.ResponseWriter, r *http.Request, queueId uuid.UUID) {
	panic("unimplemented")
}

// GetQueueSkills implements ServerInterface.
func (h QueueHandler) GetQueueSkills(w http.ResponseWriter, r *http.Request, queueId uuid.UUID) {
	panic("unimplemented")
}

// GetQueueSkill implements ServerInterface.
func (h QueueHandler) GetQueueSkill(w http.ResponseWriter, r *http.Request, queueId uuid.UUID, skillId uuid.UUID) {
	panic("unimplemented")
}

// DeleteQueueSkill implements ServerInterface.
func (h QueueHandler) DeleteQueueSkill(w http.ResponseWriter, r *http.Request, queueId uuid.UUID, skillId uuid.UUID) {
	panic("unimplemented")
}

// GetQueueInteractions implements ServerInterface.
func (h QueueHandler) GetQueueInteractions(w http.ResponseWriter, r *http.Request, queueId uuid.UUID) {
	panic("unimplemented")
}
