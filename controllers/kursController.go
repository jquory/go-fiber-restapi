package kurscontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jquory/go-fiber-restapi/models"
	// "gorm.io/gorm"
)

func GetAllKurs(ctx *fiber.Ctx) error {
	var kurs []models.Kurs
	models.DB.Find(&kurs)
	return ctx.JSON(kurs)
}

func CreateKurs(ctx *fiber.Ctx) error {
	var kurs models.Kurs
	if err := ctx.BodyParser(&kurs)
	err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	kurs.Id = uuid.New()

	if err := models.DB.Create(&kurs).Error
	err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"id": kurs.Id,
	})
}

