package v1

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nilssonr/agentside/queue"
)

// CreateQueue implements ServerInterface.
func (ah AgentsideHandler) CreateQueue(w http.ResponseWriter, r *http.Request) {
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

	created, err := ah.QueueService.CreateQueue(r.Context(), &request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, created)
}

// GetQueues implements ServerInterface.
func (ah AgentsideHandler) GetQueues(w http.ResponseWriter, r *http.Request) {
	result, err := ah.QueueService.GetQueues(r.Context(), tenantID(r))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetQueue implements ServerInterface.
func (ah AgentsideHandler) GetQueue(w http.ResponseWriter, r *http.Request, queueId uuid.UUID) {
	result, err := ah.QueueService.GetQueue(r.Context(), tenantID(r), queueId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// UpdateQueue implements ServerInterface.
func (ah AgentsideHandler) UpdateQueue(w http.ResponseWriter, r *http.Request, queueId uuid.UUID) {
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

	result, err := ah.QueueService.UpdateQueue(r.Context(), &request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// DeleteQueue implements ServerInterface.
func (ah AgentsideHandler) DeleteQueue(w http.ResponseWriter, r *http.Request, queueId uuid.UUID) {
	if err := ah.QueueService.DeleteQueue(r.Context(), tenantID(r), queueId.String()); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}

// UpsertQueueSkill implements ServerInterface.
func (ah AgentsideHandler) UpsertQueueSkill(w http.ResponseWriter, r *http.Request, queueId uuid.UUID) {
	panic("unimplemented")
}

// GetQueueSkills implements ServerInterface.
func (ah AgentsideHandler) GetQueueSkills(w http.ResponseWriter, r *http.Request, queueId uuid.UUID) {
	panic("unimplemented")
}

// GetQueueSkill implements ServerInterface.
func (ah AgentsideHandler) GetQueueSkill(w http.ResponseWriter, r *http.Request, queueId uuid.UUID, skillId uuid.UUID) {
	panic("unimplemented")
}

// DeleteQueueSkill implements ServerInterface.
func (ah AgentsideHandler) DeleteQueueSkill(w http.ResponseWriter, r *http.Request, queueId uuid.UUID, skillId uuid.UUID) {
	panic("unimplemented")
}

// GetQueueInteractions implements ServerInterface.
func (ah AgentsideHandler) GetQueueInteractions(w http.ResponseWriter, r *http.Request, queueId uuid.UUID) {
	panic("unimplemented")
}
