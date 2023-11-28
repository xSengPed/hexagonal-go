package routes

import (
	postcore "hexa-go/pkg/v1/core/post"
	"hexa-go/pkg/v1/core/post/repository"
	postapi "hexa-go/pkg/v1/post"

	"github.com/gofiber/fiber/v2"
)

// ? Add Router Config Interface Here !

type routes struct {
}

func NewRoute() *routes {
	return &routes{}
}

func (r *routes) InitializeRouter() *fiber.App {
	router := newFiber()

	v1 := router.Group("/api/v1")
	postGroup := v1.Group("/post")

	postDB := repository.NewPostRepositoryDB()
	postService := postcore.NewService(postDB)
	createPostHandler := postapi.NewCreatePostHandler(postService.CreatePostFunc)
	getPostHandler := postapi.NewGetPostHandler(postService.GetPostFunc)
	postGroup.Post("/create", createPostHandler.CreatePostHandler)
	postGroup.Get("/get", getPostHandler.GetPost)

	return router
}
func newFiber() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app
}
