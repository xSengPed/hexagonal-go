package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func LogTest(c *fiber.Ctx) error {
	id := c.Query("id")
	c.Request().Header.Add("token", "12314123")

	if id == "400" {
		return errors.New("id error")
	}

	return c.Next() // Do next Func
}
