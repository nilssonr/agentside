package skill

import "time"

type Skill struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	TenantID       string    `json:"tenantId"`
	LastModifiedBy string    `json:"lastModifiedBy"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
}
