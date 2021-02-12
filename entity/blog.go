package entity

import (
	"time"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

//Blog represents blogs table in database
type Blog struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Slug      string         `gorm:"type:varchar(254)" json:"slug"`
	Title     string         `gorm:"type:varchar(254)" json:"title"`
	Thumbnail string         `gorm:"type:varchar(254)" json:"thumbnail"`
	AuthorID  uuid.UUID      `gorm:"not null" json:"-"`
	Author    User           `gorm:"foreignkey:AuthorID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"author"`
	Content   string         `gorm:"text" json:"content"`
	Tags      pq.StringArray `gorm:"type:varchar(64)[]"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt *time.Time     `sql:"index" json:"deleted_at"`
}
