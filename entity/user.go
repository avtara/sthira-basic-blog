package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//User represents users table in database
type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name      string     `gorm:"type:varchar(255)" json:"name"`
	Email     string     `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password  string     `gorm:"->;<-;not null" json:"-"`
	Token     string     `gorm:"-" json:"token,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
