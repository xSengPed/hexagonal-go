package postapi

import (
	postcore "hexa-go/pkg/v1/core/post"

	"github.com/gofiber/fiber/v2"
)

type getPostHandler struct {
	getPost postcore.GetPostFunc
}

func NewGetPostHandler(getPost postcore.GetPostFunc) *getPostHandler {
	return &getPostHandler{getPost}
}

func (h *getPostHandler) GetPost(c *fiber.Ctx) error {
	results, err := h.getPost()

	if err != nil {
		return c.SendStatus(500)
	}
	return c.Status(200).JSON(results)
}
