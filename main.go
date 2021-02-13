package main

import (
	"log"
	"time"

	"github.com/avtara/sthira-simple-blog/config"
	"github.com/avtara/sthira-simple-blog/controller"
	"github.com/avtara/sthira-simple-blog/middleware"
	"github.com/avtara/sthira-simple-blog/repository"
	"github.com/avtara/sthira-simple-blog/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"gorm.io/gorm"
)

func main() {

	var (
		db                  *gorm.DB                       = config.SetupDatabaseConnection()
		userRepository      repository.UserRepository      = repository.NewUserRepository(db)
		blogRepository      repository.BlogRepository      = repository.NewBlogRepository(db)
		contactUsRepository repository.ContactUsRepository = repository.NewContactRepository(db)
		jwtService          service.JWTService             = service.NewJWTService()
		authService         service.AuthService            = service.NewAuthService(userRepository)
		blogService         service.BlogService            = service.NewBlogService(blogRepository)
		contactUsService    service.ContactUsService       = service.NewContactUsService(contactUsRepository)
		authController      controller.AuthController      = controller.NewAuthController(authService, jwtService)
		blogController      controller.BlogController      = controller.NewBlogController(blogService)
		contactUsController controller.ContactUsController = controller.NewContactUsController(contactUsService)
	)

	defer config.CloseDatabaseConnection(db)

	r := gin.Default()
	store, err := redis.NewStore(10, "tcp", "178.128.25.31:6379", "secret123", []byte("secret"))
	if err != nil {
		log.Fatalln(err)
	}
	store.Options(sessions.Options{MaxAge: 3600 * 24 * 30 * 12})
	r.Use(sessions.Sessions("mysession", store))

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
		authRoutes.GET("/verify", authController.ValidateToken, middleware.AuthorizeJWT(jwtService))
	}

	blogRoutes := r.Group("api/stories")
	{

		blogRoutes.GET("/", blogController.All)
		blogRoutes.GET("/:tag", blogController.BlogByTag)
	}

	articleRoutes := r.Group("api/article")
	{

		articleRoutes.GET("/:slug", blogController.BlogBySlug)
	}

	contacUsRoutes := r.Group("api/contact")
	{

		contacUsRoutes.POST("/", contactUsController.ContactUs)
	}

	r.Run(":8081")
}
