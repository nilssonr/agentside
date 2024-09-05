package v1

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nilssonr/agentside/internal/queue"
	"github.com/nilssonr/agentside/internal/skill"
	"github.com/nilssonr/agentside/internal/tenant"
	"github.com/nilssonr/agentside/internal/user"
)

type AgentsideHandler struct {
	TenantService    tenant.Service
	UserService      user.Service
	UserSkillService user.SkillService
	SkillService     skill.Service
	QueueService     queue.Service
}

// CreateCustomer implements ServerInterface.
func (ah AgentsideHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// CreateCustomerEmailAddress implements ServerInterface.
func (ah AgentsideHandler) CreateCustomerEmailAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	panic("unimplemented")
}

// CreateCustomerPhoneNumber implements ServerInterface.
func (ah AgentsideHandler) CreateCustomerPhoneNumber(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	panic("unimplemented")
}

// CreateInteraction implements ServerInterface.
func (ah AgentsideHandler) CreateInteraction(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// DeleteCustomer implements ServerInterface.
func (ah AgentsideHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	panic("unimplemented")
}

// DeleteCustomerEmailAddress implements ServerInterface.
func (ah AgentsideHandler) DeleteCustomerEmailAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, emailAddressId uuid.UUID) {
	panic("unimplemented")
}

// DeleteCustomerPhoneNumber implements ServerInterface.
func (ah AgentsideHandler) DeleteCustomerPhoneNumber(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, phoneNumberId uuid.UUID) {
	panic("unimplemented")
}

// DeleteQueueSkill implements ServerInterface.
func (ah AgentsideHandler) DeleteQueueSkill(w http.ResponseWriter, r *http.Request, queueId uuid.UUID, skillId uuid.UUID) {
	panic("unimplemented")
}

// GetCustomer implements ServerInterface.
func (ah AgentsideHandler) GetCustomer(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	panic("unimplemented")
}

// GetCustomerEmailAddress implements ServerInterface.
func (ah AgentsideHandler) GetCustomerEmailAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, emailAddressId uuid.UUID) {
	panic("unimplemented")
}

// GetCustomerEmailAddresses implements ServerInterface.
func (ah AgentsideHandler) GetCustomerEmailAddresses(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	panic("unimplemented")
}

// GetCustomerPhoneNumber implements ServerInterface.
func (ah AgentsideHandler) GetCustomerPhoneNumber(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, phoneNumberId uuid.UUID) {
	panic("unimplemented")
}

// GetCustomerPhoneNumbers implements ServerInterface.
func (ah AgentsideHandler) GetCustomerPhoneNumbers(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	panic("unimplemented")
}

// GetInteraction implements ServerInterface.
func (ah AgentsideHandler) GetInteraction(w http.ResponseWriter, r *http.Request, interactionId uuid.UUID) {
	panic("unimplemented")
}

// GetInteractions implements ServerInterface.
func (ah AgentsideHandler) GetInteractions(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetQueueInteractions implements ServerInterface.
func (ah AgentsideHandler) GetQueueInteractions(w http.ResponseWriter, r *http.Request, queueId uuid.UUID) {
	panic("unimplemented")
}

// GetQueueSkill implements ServerInterface.
func (ah AgentsideHandler) GetQueueSkill(w http.ResponseWriter, r *http.Request, queueId uuid.UUID, skillId uuid.UUID) {
	panic("unimplemented")
}

// GetQueueSkills implements ServerInterface.
func (ah AgentsideHandler) GetQueueSkills(w http.ResponseWriter, r *http.Request, queueId uuid.UUID) {
	panic("unimplemented")
}

// UpdateCustomer implements ServerInterface.
func (ah AgentsideHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request, customerId uuid.UUID) {
	panic("unimplemented")
}

// UpdateCustomerEmailAddress implements ServerInterface.
func (ah AgentsideHandler) UpdateCustomerEmailAddress(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, emailAddressId uuid.UUID) {
	panic("unimplemented")
}

// UpdateCustomerPhoneNumber implements ServerInterface.
func (ah AgentsideHandler) UpdateCustomerPhoneNumber(w http.ResponseWriter, r *http.Request, customerId uuid.UUID, phoneNumberId uuid.UUID) {
	panic("unimplemented")
}

// UpsertQueueSkill implements ServerInterface.
func (ah AgentsideHandler) UpsertQueueSkill(w http.ResponseWriter, r *http.Request, queueId uuid.UUID) {
	panic("unimplemented")
}

var _ ServerInterface = (*AgentsideHandler)(nil)

// CreateTenant implements ServerInterface.
func (ah AgentsideHandler) CreateTenant(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetTenants implements ServerInterface.
func (ah AgentsideHandler) GetTenants(w http.ResponseWriter, r *http.Request) {
	t, err := ah.TenantService.GetTenants(r.Context())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, t)
}

// GetTenant implements ServerInterface.
func (ah AgentsideHandler) GetTenant(w http.ResponseWriter, r *http.Request, tenantId uuid.UUID) {
	t, err := ah.TenantService.GetTenant(r.Context(), tenantId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, t)
}

// UpdateTenant implements ServerInterface.
func (ah AgentsideHandler) UpdateTenant(w http.ResponseWriter, r *http.Request, tenantId uuid.UUID) {
	panic("unimplemented")
}

// DeleteTenant implements ServerInterface.
func (ah AgentsideHandler) DeleteTenant(w http.ResponseWriter, r *http.Request, tenantId uuid.UUID) {
	panic("unimplemented")
}

// CreateUser implements ServerInterface.
func (ah AgentsideHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	created, err := ah.UserService.CreateUser(r.Context(), arg)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, created)
}

// GetUsers implements ServerInterface.
func (ah AgentsideHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := ah.UserService.GetUsers(r.Context(), tenantID(r))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, users)
}

// GetUser implements ServerInterface.
func (ah AgentsideHandler) GetUser(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
	usr, err := ah.UserService.GetUser(r.Context(), tenantID(r), userId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, usr)
}

// UpdateUser implements ServerInterface.
func (ah AgentsideHandler) UpdateUser(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
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

	updated, err := ah.UserService.UpdateUser(r.Context(), arg)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, updated)
}

// DeleteUser implements ServerInterface.
func (ah AgentsideHandler) DeleteUser(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
	err := ah.UserService.DeleteUser(r.Context(), tenantID(r), userId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}

// CreateSkill implements ServerInterface.
func (ah AgentsideHandler) CreateSkill(w http.ResponseWriter, r *http.Request) {
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

	result, err := ah.SkillService.CreateSkill(r.Context(), request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetSkills implements ServerInterface.
func (ah AgentsideHandler) GetSkills(w http.ResponseWriter, r *http.Request) {
	result, err := ah.SkillService.GetSkills(r.Context(), tenantID(r))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetSkill implements ServerInterface.
func (ah AgentsideHandler) GetSkill(w http.ResponseWriter, r *http.Request, skillId uuid.UUID) {
	result, err := ah.SkillService.GetSkill(r.Context(), tenantID(r), skillId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// UpdateSkill implements ServerInterface.
func (ah AgentsideHandler) UpdateSkill(w http.ResponseWriter, r *http.Request, skillId uuid.UUID) {
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

	result, err := ah.SkillService.UpdateSkill(r.Context(), request)
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// DeleteSkill implements ServerInterface.
func (ah AgentsideHandler) DeleteSkill(w http.ResponseWriter, r *http.Request, skillId uuid.UUID) {
	if err := ah.SkillService.DeleteSkill(r.Context(), tenantID(r), skillId.String()); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}

// UpsertUserSkill implements ServerInterface.
func (ah AgentsideHandler) UpsertUserSkill(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
	var body UpsertUserSkillRequest
	if err := render.DecodeJSON(r.Body, &body); err != nil {
		handleError(w, r, err)
		return
	}
	defer r.Body.Close()

	upserted, err := ah.UserSkillService.UpsertSkill(
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
func (ah AgentsideHandler) GetUserSkills(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
	result, err := ah.UserSkillService.GetSkills(r.Context(), userID(r))
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// GetUserSkill implements ServerInterface.
func (ah AgentsideHandler) GetUserSkill(w http.ResponseWriter, r *http.Request, userId uuid.UUID, skillId uuid.UUID) {
	result, err := ah.UserSkillService.GetSkill(r.Context(), userId.String(), skillId.String())
	if err != nil {
		handleError(w, r, err)
		return
	}

	render.JSON(w, r, result)
}

// DeleteUserSkill implements ServerInterface.
func (ah AgentsideHandler) DeleteUserSkill(w http.ResponseWriter, r *http.Request, userId uuid.UUID, skillId uuid.UUID) {
	if err := ah.UserSkillService.DeleteSkill(r.Context(), userId.String(), skillId.String()); err != nil {
		handleError(w, r, err)
		return
	}

	render.NoContent(w, r)
}

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

func tenantID(r *http.Request) string {
	tenantID, ok := r.Context().Value("tenant_id").(string)
	if !ok {
		return ""
	}

	return tenantID
}

func userID(r *http.Request) string {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		return ""
	}

	return userID
}
