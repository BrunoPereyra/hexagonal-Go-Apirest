package middleware

import (
	"hexagonal/pkg/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UseExtractor() fiber.Handler {
	return func(c *fiber.Ctx) error {
		autooHeader := c.Get("Authorization")
		token := strings.Replace(autooHeader, "Bearer ", "", 1)
		_, _id, err := auth.ExtractDataFromToken(token)
		if err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
				"Message": "Unauthorized",
			})
		}
		c.Context().SetUserValue("_id", _id)
		return c.Next()
	}
}
