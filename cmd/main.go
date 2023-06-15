package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"hexagonal/internal/middleware"
	deliveryPost "hexagonal/internal/post/delivery"
	deliveryUser "hexagonal/internal/user/delivery"
)

func main() {
	app := fiber.New()

	app.Post("/signup", deliveryUser.Signup)
	app.Post("/login", deliveryUser.Login)

	app.Post("/UploadPost", middleware.UseExtractor(), deliveryPost.CreatePost)

	log.Fatal(app.Listen(":3001"))
}
