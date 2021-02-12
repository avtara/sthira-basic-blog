package controller

import (
	"fmt"
	"net/http"

	"github.com/avtara/sthira-simple-blog/entity"
	"github.com/avtara/sthira-simple-blog/helper"
	"github.com/avtara/sthira-simple-blog/service"
	"github.com/gin-gonic/gin"
)

//BlogController is a contract about something that service can do
type BlogController interface {
	All(ctx *gin.Context)
	BlogByTag(ctx *gin.Context)
	BlogBySlug(ctx *gin.Context)
}

type blogController struct {
	blogService service.BlogService
}

//NewBlogController create a new instances of BoookController
func NewBlogController(blogService service.BlogService) BlogController {
	return &blogController{
		blogService: blogService,
	}
}

func (c *blogController) All(ctx *gin.Context) {
	var blogs []entity.Blog = c.blogService.All()
	res := helper.BuildResponse(true, "OK", blogs)
	ctx.JSON(http.StatusOK, res)
}

func (c *blogController) BlogByTag(ctx *gin.Context) {
	tag := ctx.Param("tag")
	var blog []entity.Blog = c.blogService.BlogByTag(tag)
	res := helper.BuildResponse(true, "OK", blog)
	ctx.JSON(http.StatusOK, res)
}

func (c *blogController) BlogBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	var blog entity.Blog = c.blogService.BlogBySlug(slug)
	fmt.Println(blog)
	res := helper.BuildResponse(true, "OK", blog)
	ctx.JSON(http.StatusOK, res)
}
