package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func registerSkillRouter(r chi.Router) {
	r.Route("/skills", func(r chi.Router) {
		r.Post("/", createSkill)
		r.Get("/", getSkills)
		r.Route("/{skillID}", func(r chi.Router) {
			r.Get("/", getSkill)
			r.Put("/", updateSkill)
			r.Delete("/", deleteSkill)
		})
	})
}

type createSkillRequest struct {
	Name string `json:"name"`
}

type updateSkillRequest struct {
	Name string `json:"name"`
}

func createSkill(w http.ResponseWriter, r *http.Request) {

}

func getSkills(w http.ResponseWriter, r *http.Request) {

}

func getSkill(w http.ResponseWriter, r *http.Request) {

}

func updateSkill(w http.ResponseWriter, r *http.Request) {

}

func deleteSkill(w http.ResponseWriter, r *http.Request) {

}
