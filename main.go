package main

import (
	"github.com/gofiber/fiber/v2"
	kurscontrollers "github.com/jquory/go-fiber-restapi/controllers"
	"github.com/jquory/go-fiber-restapi/middleware"
	"github.com/jquory/go-fiber-restapi/models"
)

func main() {
	models.InitialDatabase()

	app := fiber.New()
	app.Use(middleware.AuthMiddleware)
	kurs := app.Group("/api")

	kurs.Get("/kurs", kurscontrollers.GetAllKurs)

	app.Listen(":8080")
}