package v1

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nilssonr/agentside/skill"
)

type SkillHandler struct {
	SkillService skill.Service
}

func NewSkillHandler(ss skill.Service) SkillHandler {
	return SkillHandler{
		SkillService: ss,
	}
}

// CreateSkill implements ServerInterface.
func (h SkillHandler) CreateSkill(w http.ResponseWriter, r *http.Request) {
	var body CreateSkillRequest
	if err := render.DecodeJSON(r.Body, &body); err != nil {
		handleError(w, r, err)
		return
	}
	defer r.Body.Close()

	request := &skill.Skill{
		Name:           body.Name,
		TenantID:       tenantID(r),
		LastModifiedBy: userID(r),
		LastModifiedAt: time.Now(),
	}

	result, err := h.SkillService.CreateSkill(r.Context(), request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetSkills implements ServerInterface.
func (h SkillHandler) GetSkills(w http.ResponseWriter, r *http.Request) {
	result, err := h.SkillService.GetSkills(r.Context(), tenantID(r))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetSkill implements ServerInterface.
func (h SkillHandler) GetSkill(w http.ResponseWriter, r *http.Request, skillId uuid.UUID) {
	result, err := h.SkillService.GetSkill(r.Context(), tenantID(r), skillId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// UpdateSkill implements ServerInterface.
func (h SkillHandler) UpdateSkill(w http.ResponseWriter, r *http.Request, skillId uuid.UUID) {
	var body UpdateSkillRequest
	if err := render.DecodeJSON(r.Body, &body); err != nil {
		handleError(w, r, err)
		return
	}
	defer r.Body.Close()

	request := &skill.Skill{
		ID:             skillId.String(),
		Name:           body.Name,
		TenantID:       tenantID(r),
		LastModifiedBy: userID(r),
		LastModifiedAt: time.Now(),
	}

	result, err := h.SkillService.UpdateSkill(r.Context(), request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// DeleteSkill implements ServerInterface.
func (h SkillHandler) DeleteSkill(w http.ResponseWriter, r *http.Request, skillId uuid.UUID) {
	if err := h.SkillService.DeleteSkill(r.Context(), tenantID(r), skillId.String()); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}
