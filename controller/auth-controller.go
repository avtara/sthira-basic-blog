package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/avtara/sthira-simple-blog/dto"
	"github.com/avtara/sthira-simple-blog/entity"
	"github.com/avtara/sthira-simple-blog/helper"
	"github.com/avtara/sthira-simple-blog/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

//AuthController interface is a contract what controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	ValidateToken(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	session := sessions.Default(ctx)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		if session.Get(v.Email) == nil {
			generatedToken := c.jwtService.GenerateToken(v.ID)
			session.Set(v.Email, generatedToken)
			session.Save()
			v.Token = fmt.Sprint(session.Get(v.Email))
		} else {
			v.Token = fmt.Sprint(session.Get(v.Email))
		}

		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(createdUser.ID)
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *authController) ValidateToken(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	splitToken := strings.Split(authHeader, "Bearer ")
	token, err := c.jwtService.ValidateToken(splitToken[1])
	if !token.Valid {
		response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	} else {
		userID := c.getUserIDByToken(splitToken[1])
		convertUUID, err := uuid.FromString(userID)
		if err == nil {
			authResult := c.authService.FindByID(convertUUID)
			authResult.Token = splitToken[1]
			response := helper.BuildResponse(true, "OK!", authResult)
			ctx.JSON(http.StatusOK, response)
			return
		}
	}
}

func (c *authController) getUserIDByToken(token string) string {
	fmt.Println(token)
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
