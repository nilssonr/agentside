package v1

import (
	"net/http"

	"github.com/nilssonr/agentside/customer"
	"github.com/nilssonr/agentside/queue"
	"github.com/nilssonr/agentside/skill"
	"github.com/nilssonr/agentside/tenant"
	"github.com/nilssonr/agentside/user"
)

type AgentsideHandler struct {
	TenantService               tenant.Service
	UserService                 user.Service
	UserSkillService            user.SkillService
	SkillService                skill.Service
	QueueService                queue.Service
	CustomerService             customer.Service
	CustomerPhoneNumberService  customer.PhoneNumberService
	CustomerEmailAddressService customer.EmailAddressService
	CustomerAddressService      customer.AddressService
	CustomerNoteService         customer.NoteService
}

var _ ServerInterface = (*AgentsideHandler)(nil)

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
