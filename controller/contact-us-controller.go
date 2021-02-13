package controller

import (
	"net/http"

	"github.com/avtara/sthira-simple-blog/dto"
	"github.com/avtara/sthira-simple-blog/helper"
	"github.com/avtara/sthira-simple-blog/service"
	"github.com/gin-gonic/gin"
)

//ContactUsController interface is a contract what controller can do
type ContactUsController interface {
	ContactUs(ctx *gin.Context)
}

type contactUsController struct {
	contactUsService service.ContactUsService
}

//NewContactUsController creates a new instance of AuthController
func NewContactUsController(contactUsService service.ContactUsService) ContactUsController {
	return &contactUsController{
		contactUsService: contactUsService,
	}
}

func (c *contactUsController) ContactUs(ctx *gin.Context) {
	var contactUsDTO dto.ContactUsDTO
	errDTO := ctx.ShouldBind(&contactUsDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	createdUser := c.contactUsService.CreateMessage(contactUsDTO)
	response := helper.BuildResponse(true, "Successfully sent us a message, please wait for our reply 1x23.5 hours", createdUser)
	ctx.JSON(http.StatusCreated, response)
}
