package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//ContactUs represents blogs table in database
type ContactUs struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name      string     `gorm:"type:varchar(254)" json:"name"`
	Email     string     `gorm:"type:varchar(254)" json:"email"`
	Message   string     `gorm:"type:varchar(254)" json:"message"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
