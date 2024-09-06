package queue

import "time"

type Queue struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	TenantID       string    `json:"tenantId"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	LastModifiedBy string    `json:"lastModifiedBy"`
}
