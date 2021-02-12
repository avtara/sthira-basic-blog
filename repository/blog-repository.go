package repository

import (
	"github.com/avtara/sthira-simple-blog/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//BlogRepository is contract what userRepository can do to db
type BlogRepository interface {
	AllBlogs() []entity.Blog
	BlogByTags(tags string) []entity.Blog
	BlogBySlug(slug string) entity.Blog
}

type blogConnection struct {
	connection *gorm.DB
}

//NewBlogRepository is creates a new instance of UserRepository
func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &blogConnection{
		connection: db,
	}
}

func (db *blogConnection) AllBlogs() []entity.Blog {
	var blogs []entity.Blog
	db.connection.Preload(clause.Associations).Find(&blogs)
	return blogs
}

func (db *blogConnection) BlogByTags(tags string) []entity.Blog {
	var blogs []entity.Blog
	db.connection.Preload(clause.Associations).Where("tags[1] = ?", tags).Or("tags[2] = ?", tags).Or("tags[3] = ?", tags).Order("created_at asc").Find(&blogs)
	return blogs
}

func (db *blogConnection) BlogBySlug(slug string) entity.Blog {
	var blogs entity.Blog
	db.connection.Preload(clause.Associations).Find(&blogs, "slug = ?", slug)
	return blogs
}
