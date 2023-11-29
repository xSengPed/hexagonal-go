package routes

import (
	"hexa-go/configs"
	commentapi "hexa-go/pkg/v1/comment"
	commentcore "hexa-go/pkg/v1/core/comment"
	commentrepo "hexa-go/pkg/v1/core/comment/repository"
	postcore "hexa-go/pkg/v1/core/post"
	"hexa-go/pkg/v1/core/post/repository"
	postapi "hexa-go/pkg/v1/post"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// ? Add Router Config Interface Here !

type routes struct {
	mongoClient *mongo.Client
	cfg         *configs.Config
}

func NewRoute(mongoClient *mongo.Client, cfg *configs.Config) *routes {
	return &routes{mongoClient, cfg}
}

func (r *routes) InitializeRouter() *fiber.App {
	mongodb := r.mongoClient.Database(r.cfg.MongoConfig.DBName)
	router := newFiber()

	v1 := router.Group("/api/v1")

	postGroup := v1.Group("/post")
	postDB := repository.NewPostRepositoryDB(mongodb.Collection(r.cfg.AppConstants.PostCollectionName))
	postService := postcore.NewService(postDB)
	createPostHandler := postapi.NewCreatePostHandler(postService.CreatePostFunc)
	getPostHandler := postapi.NewGetPostHandler(postService.GetPostFunc)
	postGroup.Post("/create", createPostHandler.CreatePostHandler)
	postGroup.Get("/get", getPostHandler.GetPost)

	commentGroup := v1.Group("/comment")
	commentDB := commentrepo.NewCommentRepositoryDB(mongodb.Collection(r.cfg.AppConstants.CommentCollectionName))
	commentService := commentcore.NewService(commentDB)
	createCommentHandler := commentapi.NewCreateCommentHandler(commentService.CreateComment)
	commentGroup.Post("/create", createCommentHandler.CreateComment)

	return router
}
func newFiber() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app
}
