package v1

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nilssonr/agentside/user"
)

type UserHandler struct {
	UserService      user.Service
	UserSkillService user.SkillService
}

func NewUserHandler(
	userService user.Service,
	userSkillService user.SkillService,
) UserHandler {
	return UserHandler{
		UserService:      userService,
		UserSkillService: userSkillService,
	}
}

// CreateUser implements ServerInterface.
func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body CreateUserRequest
	if err := render.DecodeJSON(r.Body, &body); err != nil {
		handleError(w, r, err)
		return
	}
	defer r.Body.Close()

	arg := &user.User{
		Firstname:      body.FirstName,
		Lastname:       body.LastName,
		EmailAddress:   string(body.EmailAddress),
		TenantID:       tenantID(r),
		LastModifiedAt: time.Now(),
		LastModifiedBy: userID(r),
	}

	created, err := h.UserService.CreateUser(r.Context(), arg)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, created)
}

// GetUsers implements ServerInterface.
func (h UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.GetUsers(r.Context(), tenantID(r))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, users)
}

// GetUser implements ServerInterface.
func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
	usr, err := h.UserService.GetUser(r.Context(), tenantID(r), userId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, usr)
}

// UpdateUser implements ServerInterface.
func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
	var body UpdateUserRequest
	if err := render.DecodeJSON(r.Body, &body); err != nil {
		handleError(w, r, err)
		return
	}
	defer r.Body.Close()

	arg := &user.User{
		ID:             userId.String(),
		Firstname:      body.FirstName,
		Lastname:       body.LastName,
		LastModifiedAt: time.Now(),
		LastModifiedBy: userID(r),
		TenantID:       tenantID(r),
	}

	updated, err := h.UserService.UpdateUser(r.Context(), arg)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, updated)
}

// DeleteUser implements ServerInterface.
func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
	err := h.UserService.DeleteUser(r.Context(), tenantID(r), userId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}

// UpsertUserSkill implements ServerInterface.
func (h UserHandler) UpsertUserSkill(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
	var body UpsertUserSkillRequest
	if err := render.DecodeJSON(r.Body, &body); err != nil {
		handleError(w, r, err)
		return
	}
	defer r.Body.Close()

	upserted, err := h.UserSkillService.UpsertSkill(
		r.Context(),
		userId.String(),
		body.SkillId.String(),
		int(body.Level),
	)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, upserted)
}

// GetUserSkills implements ServerInterface.
func (h UserHandler) GetUserSkills(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
	result, err := h.UserSkillService.GetSkills(r.Context(), userID(r))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetUserSkill implements ServerInterface.
func (h UserHandler) GetUserSkill(w http.ResponseWriter, r *http.Request, userId uuid.UUID, skillId uuid.UUID) {
	result, err := h.UserSkillService.GetSkill(r.Context(), userId.String(), skillId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// DeleteUserSkill implements ServerInterface.
func (h UserHandler) DeleteUserSkill(w http.ResponseWriter, r *http.Request, userId uuid.UUID, skillId uuid.UUID) {
	if err := h.UserSkillService.DeleteSkill(r.Context(), userId.String(), skillId.String()); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}
