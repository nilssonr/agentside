package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/nilssonr/agentside/user"
)

type userHandler struct {
	UserService      user.Service
	UserSkillService user.SkillService
}

func (h userHandler) Register(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.createUser)
		r.Get("/", h.getUsers)

		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", h.getUser)
			r.Put("/", h.updateUser)
			r.Delete("/", h.deleteUser)

			r.Route("/skills", func(r chi.Router) {
				r.Put("/", h.upsertUserSkill)
				r.Get("/", h.getUserSkills)

				r.Route("/{skillID}", func(r chi.Router) {
					r.Get("/", h.getUserSkill)
					r.Delete("/", h.deleteUserSkill)
				})
			})
		})
	})
}

type createUserRequest struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
}

type updateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (h userHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var body createUserRequest
	if err := render.Decode(r, &body); err != nil {
		handleError(w, r, err)
		return
	}

	request := &user.User{
		Firstname:      body.FirstName,
		Lastname:       body.LastName,
		EmailAddress:   body.EmailAddress,
		TenantID:       tenantID(r),
		LastModifiedAt: time.Now(),
		LastModifiedBy: userID(r),
	}

	result, err := h.UserService.CreateUser(r.Context(), request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h userHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	result, err := h.UserService.GetUsers(r.Context(), tenantID(r))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h userHandler) getUser(w http.ResponseWriter, r *http.Request) {
	result, err := h.UserService.GetUser(r.Context(), tenantID(r), chi.URLParam(r, "userID"))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h userHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	var body updateUserRequest
	if err := render.Decode(r, &body); err != nil {
		handleError(w, r, err)
		return
	}

	request := &user.User{
		ID:             chi.URLParam(r, "userID"),
		Firstname:      body.FirstName,
		Lastname:       body.LastName,
		TenantID:       tenantID(r),
		LastModifiedAt: time.Now(),
		LastModifiedBy: userID(r),
	}

	result, err := h.UserService.UpdateUser(r.Context(), request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

func (h userHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	if err := h.UserService.DeleteUser(r.Context(), tenantID(r), chi.URLParam(r, "userID")); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}

func (h userHandler) upsertUserSkill(w http.ResponseWriter, r *http.Request) {

}

func (h userHandler) getUserSkills(w http.ResponseWriter, r *http.Request) {

}

func (h userHandler) getUserSkill(w http.ResponseWriter, r *http.Request) {

}

func (h userHandler) deleteUserSkill(w http.ResponseWriter, r *http.Request) {

}
