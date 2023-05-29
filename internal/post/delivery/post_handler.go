package delivery

import (
	"hexagonal/internal/post/application"
	"hexagonal/internal/post/domain"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("imgPostUrl")
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad request",
		})
	}
	var Post domain.Post
	if errParser := c.BodyParser(&Post); errParser != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad request",
		})
	}
	IDUserContext := c.Context().UserValue("_id")
	IDUserString, stringIDUserok := IDUserContext.(string)
	if !stringIDUserok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "StatusInternalServerError",
		})
	}

	IDUserPrimitive, _ := primitive.ObjectIDFromHex(IDUserString)
	Post.UserID = IDUserPrimitive

	ResCreatePost, resPost, ErrCreatePost := application.CreatePost(fileHeader, &Post)
	if ErrCreatePost != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": ErrCreatePost.Error(),
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": ResCreatePost,
			"data":    resPost,
		})
	}
}
