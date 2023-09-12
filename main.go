package main

import (
	"github.com/gofiber/fiber/v2"
	kurscontrollers "github.com/jquory/go-fiber-restapi/controllers"
	"github.com/jquory/go-fiber-restapi/models"
)

func authMiddleware(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if (authHeader == "") {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return ctx.Next()
}

func main() {
	models.InitialDatabase()

	app := fiber.New()
	app.Use(authMiddleware)
	kurs := app.Group("/api")

	kurs.Get("/kurs", kurscontrollers.GetAllKurs)

	app.Listen(":8080")
}