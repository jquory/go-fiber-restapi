package kurscontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jquory/go-fiber-restapi/models"
)

func GetAllKurs(ctx *fiber.Ctx) error {
	var kurs []models.Kurs
	models.DB.Find(&kurs)
	return ctx.JSON(kurs)
}

// func CreateKurs(ctx *fiber.Ctx) error {
// 	var kurs models.Kurs
// 	if err := 

// }