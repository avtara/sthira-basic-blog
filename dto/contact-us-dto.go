package dto

//ContactUsDTO is used when client post from /contact-us url
type ContactUsDTO struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Email   string `json:"email" form:"email" binding:"required,email" `
	Message string `json:"message" form:"message" binding:"required"`
}
