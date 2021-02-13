package repository

import (
	"github.com/avtara/sthira-simple-blog/entity"
	"gorm.io/gorm"
)

//ContactUsRepository is contract what can do to db
type ContactUsRepository interface {
	InsertMessage(message entity.ContactUs) entity.ContactUs
}

type contactUsConnection struct {
	connection *gorm.DB
}

//NewContactRepository is creates a new instance of
func NewContactRepository(db *gorm.DB) ContactUsRepository {
	return &contactUsConnection{
		connection: db,
	}
}

func (db *contactUsConnection) InsertMessage(message entity.ContactUs) entity.ContactUs {
	db.connection.Save(&message)
	return message
}
