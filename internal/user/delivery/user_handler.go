package delivery

import (
	"hexagonal/internal/user/application"

	"github.com/gofiber/fiber/v2"
)

// UserHandler maneja las solicitudes relacionadas con los usuarios a través de HTTP.

// CreateUser maneja la solicitud de creación de un nuevo usuario.
var UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(c *fiber.Ctx) error {
	// Lógica para recibir los datos de la solicitud y llamar al servicio UserService para crear el usuario.

	// Parsear los datos de la solicitud...

	if err := c.BodyParser(&UserData); err != nil {
		return err
	}

	// Invocar el servicio UserService para crear el usuario...
	user, err := application.CreateUser(UserData.Username, UserData.Password)

	if err != nil {
		return c.JSON(err.Error())
	}

	// Enviar la respuesta adecuada...
	return c.JSON(user)
}

// Otros métodos y funcionalidad relacionados con la entrega de usuarios a través de HTTP.

func Login(c *fiber.Ctx) error {
	if err := c.BodyParser(&UserData); err != nil {
		return err
	}

	// Invocar el servicio UserService para crear el usuario...
	token, err := application.Login(UserData.Username, UserData.Password)
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
