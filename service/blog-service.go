package service

import (
	"github.com/avtara/sthira-simple-blog/entity"
	"github.com/avtara/sthira-simple-blog/repository"
)

//BlogService is a contract about something that service can do
type BlogService interface {
	All() []entity.Blog
	BlogByTag(tags string) []entity.Blog
	BlogBySlug(slug string) entity.Blog
}

type blogService struct {
	blogRepository repository.BlogRepository
}

//NewBlogService creates a new instance of AuthService
func NewBlogService(blogRepository repository.BlogRepository) BlogService {
	return &blogService{
		blogRepository: blogRepository,
	}
}

func (service *blogService) All() []entity.Blog {
	return service.blogRepository.AllBlogs()
}

func (service *blogService) BlogByTag(tags string) []entity.Blog {
	return service.blogRepository.BlogByTags(tags)
}

func (service *blogService) BlogBySlug(slug string) entity.Blog {
	return service.blogRepository.BlogBySlug(slug)
}
