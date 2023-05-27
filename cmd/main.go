package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"hexagonal/internal/user/delivery"
)

func main() {
	app := fiber.New()

	// Definir las rutas y los controladores correspondientes
	app.Post("/userCreate", delivery.CreateUser)

	// Iniciar el servidor
	log.Fatal(app.Listen(":3001"))
}
