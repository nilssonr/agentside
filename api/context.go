package api

import "net/http"

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
