package delivery

import (
	"hexagonal/internal/post/application"

	"github.com/gofiber/fiber/v2"
)

type PostModel struct {
	TextPost string `json:"Status" bson:"Status"`
}

func CreatePost(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("PostImage")
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad request",
		})
	}
	var Post PostModel
	if errParser := c.BodyParser(Post); errParser != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad request",
		})
	}
	ResCreatePost, resPost, ErrCreatePost := application.CreatePost(fileHeader, Post.TextPost)
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
