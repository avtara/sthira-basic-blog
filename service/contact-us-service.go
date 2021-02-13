package service

import (
	"log"

	"github.com/avtara/sthira-simple-blog/dto"
	"github.com/avtara/sthira-simple-blog/entity"
	"github.com/avtara/sthira-simple-blog/repository"
	"github.com/mashingan/smapping"
)

//ContactUsService is a contract about something that service can do
type ContactUsService interface {
	CreateMessage(mes dto.ContactUsDTO) entity.ContactUs
}

type contactUsService struct {
	contactUsRepository repository.ContactUsRepository
}

//NewContactUsService creates a new instance of AuthService
func NewContactUsService(contactRep repository.ContactUsRepository) ContactUsService {
	return &contactUsService{
		contactUsRepository: contactRep,
	}
}

func (service *contactUsService) CreateMessage(mes dto.ContactUsDTO) entity.ContactUs {
	messageToCreate := entity.ContactUs{}
	err := smapping.FillStruct(&messageToCreate, smapping.MapFields(&mes))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.contactUsRepository.InsertMessage(messageToCreate)
	return res
}
