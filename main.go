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
	api := app.Group("/api")
	kurs := api.Group("/kurs")

	kurs.Get("/", kurscontrollers.GetAllKurs)
	kurs.Get("/:id", kurscontrollers.ShowKurs)
	kurs.Post("/", kurscontrollers.CreateKurs)
	kurs.Put("/:id", kurscontrollers.UpdateKurs)
	kurs.Delete("/:id", kurscontrollers.DeleteKurs)

	app.Listen(":8080")
}