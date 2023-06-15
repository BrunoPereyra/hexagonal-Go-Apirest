package delivery

import (
	"hexagonal/internal/user/application"
	"hexagonal/internal/user/domain"

	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {

	fileHeader, errfileHeader := c.FormFile("avatar")
	if errfileHeader != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messages": "Bad Request",
		})
	}

	var newUser domain.UserModelValidator

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messages": "Bad Request",
		})
	}
	if err := newUser.ValidateUserFind(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	// Invocar el servicio UserService para crear el usuario...
	user, err := application.CreateUser(&newUser, fileHeader)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var UserLogin domain.LoginValidatorStruct
	if err := c.BodyParser(&UserLogin); err != nil {
		return err
	}
	errUserValidator := UserLogin.LoginValidator()
	if errUserValidator != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errUserValidator.Error(),
		})
	}

	token, err := application.Login(&UserLogin)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"token":   token,
	})
}
