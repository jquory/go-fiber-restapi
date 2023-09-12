package kurscontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jquory/go-fiber-restapi/models"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
)

func GetAllKurs(ctx *fiber.Ctx) error {
	var kurs []models.Kurs
	models.DB.Find(&kurs)
	return ctx.JSON(kurs)
}

func ShowKurs(ctx *fiber.Ctx) error {
	idStr := ctx.Params("Id")
	Id, err := uuid.Parse(idStr)
	var kurs models.Kurs
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := models.DB.First(&kurs, Id).Error
	err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(404).JSON(fiber.Map{
				"message": "Data not Found",
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
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

	kurs.Id = mssql.UniqueIdentifier(uuid.New())

	if err := models.DB.Create(&kurs).Error
	err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(kurs)
}

func UpdateKurs(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "ID not valid",
		})
	}
	var kurs models.Kurs

	if err := ctx.BodyParser(&kurs) 
	err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&kurs).RowsAffected == 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Kurs berhasil di perbarui",
	})
}

func DeleteKurs(ctx *fiber.Ctx) error {
	var idStr = ctx.Params("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var kurs models.Kurs
	if models.DB.Delete(&kurs, id).Error != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Not found",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "Data berhasil dihapus",
	})
}