package tenant

import "time"

type Tenant struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	LastModifiedBy string    `json:"lastModifiedBy"`
}
