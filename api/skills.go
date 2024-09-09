package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/nilssonr/agentside/skill"
)

type skillHandler struct {
	SkillService skill.Service
}

func (h skillHandler) Register(r chi.Router) {
	r.Route("/skills", func(r chi.Router) {
		r.Post("/", h.createSkill)
		r.Get("/", h.getSkills)

		r.Route("/{skillID}", func(r chi.Router) {
			r.Get("/", h.getSkill)
			r.Put("/", h.updateSkill)
			r.Delete("/", h.deleteSkill)
		})
	})
}

type createSkillRequest struct {
	Name string `json:"name"`
}

type updateSkillRequest struct {
	Name string `json:"name"`
}

func (h skillHandler) createSkill(w http.ResponseWriter, r *http.Request) {
	var body createSkillRequest
	if err := render.Decode(r, &body); err != nil {
		handleError(w, r, err)
		return
	}

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

func (h skillHandler) getSkills(w http.ResponseWriter, r *http.Request) {
	result, err := h.SkillService.GetSkills(r.Context(), tenantID(r))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h skillHandler) getSkill(w http.ResponseWriter, r *http.Request) {
	result, err := h.SkillService.GetSkill(r.Context(), tenantID(r), chi.URLParam(r, "skillID"))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h skillHandler) updateSkill(w http.ResponseWriter, r *http.Request) {
	var body updateSkillRequest
	if err := render.Decode(r, &body); err != nil {
		handleError(w, r, err)
		return
	}

	request := &skill.Skill{
		ID:             chi.URLParam(r, "skillID"),
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

func (h skillHandler) deleteSkill(w http.ResponseWriter, r *http.Request) {
	if err := h.SkillService.DeleteSkill(r.Context(), tenantID(r), chi.URLParam(r, "skillID")); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}
