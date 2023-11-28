package commentapi

import (
	commentcore "hexa-go/pkg/v1/core/comment"
	"hexa-go/pkg/v1/datamodels"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type createCommentHandler struct {
	createCommentFunc commentcore.CreateComment
}

func NewCreateCommentHandler(createComment commentcore.CreateComment) *createCommentHandler {
	return &createCommentHandler{createComment}
}

type createCommentReq struct {
	Desc string `json:"desc" validate:"required"`
}

func (h *createCommentHandler) CreateComment(ctx *fiber.Ctx) error {

	payload := createCommentReq{}

	err := ctx.BodyParser(&payload)
	if err != nil {
		return ctx.SendStatus(400)
	}

	err = payload.validate()
	if err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	comment := datamodels.Comment{

		Desc: payload.Desc,
	}

	err = h.createCommentFunc(ctx.Context(), comment)
	if err != nil {
		return ctx.SendStatus(500)
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"msg": "create comment successful",
	})
}

func (c *createCommentReq) validate() error {
	v := validator.New()
	return v.Struct(c)
}
