package delivery

import (
	"hexagonal/internal/user/application"

	"github.com/gofiber/fiber/v2"
)

// UserHandler maneja las solicitudes relacionadas con los usuarios a través de HTTP.

// CreateUser maneja la solicitud de creación de un nuevo usuario.
func CreateUser(c *fiber.Ctx) error {
	// Lógica para recibir los datos de la solicitud y llamar al servicio UserService para crear el usuario.

	// Parsear los datos de la solicitud...
	var userData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&userData); err != nil {
		return err
	}

	// Invocar el servicio UserService para crear el usuario...
	user, err := application.CreateUser(userData.Username, userData.Password)

	if err != nil {
		return c.JSON(err.Error())
	}

	// Enviar la respuesta adecuada...
	return c.JSON(user)
}

// Otros métodos y funcionalidad relacionados con la entrega de usuarios a través de HTTP.
