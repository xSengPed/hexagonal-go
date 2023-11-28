package postapi

import (
	postcore "hexa-go/pkg/v1/core/post"
	"hexa-go/pkg/v1/datamodels"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type createPostHandler struct {
	createPost postcore.CreatePostFunc
}

type postReq struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (v *postReq) validate() error {
	validate := validator.New()
	return validate.Struct(v)

}

type postRes struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewCreatePostHandler(createPost postcore.CreatePostFunc) *createPostHandler {
	return &createPostHandler{createPost}
}

func (h *createPostHandler) CreatePostHandler(c *fiber.Ctx) error {
	var payload postReq

	err := c.BodyParser(&payload)
	if err != nil {
		return c.SendStatus(500)
	}

	err = payload.validate()
	if err != nil {
		return c.SendStatus(400)
	}

	post := datamodels.Post{
		Title:   payload.Title,
		Content: payload.Content,
	}

	err = h.createPost(c.Context(), post)

	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"err_msg": err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"msg": "create success",
	})

}
