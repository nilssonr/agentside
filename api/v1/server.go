package v1

import (
	"net/http"
)

type ServerHandler struct {
	TenantHandler
	SkillHandler
	QueueHandler
	CustomerHandler
	InteractionHandler
	UserHandler
}

func NewServerHandler(
	ch CustomerHandler,
	ih InteractionHandler,
	qh QueueHandler,
	sh SkillHandler,
	th TenantHandler,
	uh UserHandler) ServerHandler {
	return ServerHandler{
		TenantHandler:      th,
		UserHandler:        uh,
		SkillHandler:       sh,
		QueueHandler:       qh,
		CustomerHandler:    ch,
		InteractionHandler: ih,
	}
}

var _ ServerInterface = (*ServerHandler)(nil)

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
