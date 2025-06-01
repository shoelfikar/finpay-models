package general

import "time"

type Timestamp struct {
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedBy string     `json:"created_by"`
}
