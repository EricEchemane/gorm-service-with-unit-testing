package identity

import "time"

type Role int
type Status int

const (
	Superuser Role = 1
	Admin     Role = 2
	Normal    Role = 3

	Active   Status = 1
	Inactive Status = 2
)

type User struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`

	Username string `gorm:"index:idx_username,unique" json:"username"`
	Password string `json:"password,omitempty"`
	Role     Role   `json:"role"`
	Status   Status `json:"status"`
}