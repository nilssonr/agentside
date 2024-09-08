package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
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

}

func (h skillHandler) getSkills(w http.ResponseWriter, r *http.Request) {

}

func (h skillHandler) getSkill(w http.ResponseWriter, r *http.Request) {

}

func (h skillHandler) updateSkill(w http.ResponseWriter, r *http.Request) {

}

func (h skillHandler) deleteSkill(w http.ResponseWriter, r *http.Request) {

}
