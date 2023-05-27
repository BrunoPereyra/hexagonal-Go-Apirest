package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	deliveryPost "hexagonal/internal/post/delivery"
	deliveryUser "hexagonal/internal/user/delivery"
)

func main() {
	app := fiber.New()

	// Definir las rutas y los controladores correspondientes
	app.Post("/userCreate", deliveryUser.CreateUser)
	app.Post("/login", deliveryUser.Login)

	app.Post("/CreatePost", deliveryPost.CreatePost)

	// Iniciar el servidor
	log.Fatal(app.Listen(":3001"))
}
