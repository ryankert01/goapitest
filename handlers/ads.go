package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryankert01/goapitest/database"
	"github.com/ryankert01/goapitest/models"
	"github.com/ryankert01/goapitest/utils"
	// "gorm.io/datatypes"
	// "fmt"
)

func GetAds(c *fiber.Ctx) error {
	var ads []models.AD
	database.DB.Db.Find(&ads)
	return c.Status(200).JSON(ads)
}

func CreateAd(c *fiber.Ctx) error {
	ad := new(models.AD)
	if err := c.BodyParser(ad); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if utils.IsAdAlive(*ad) {
		ad.IsActive = true
	} else {
		ad.IsActive = false
	}

	database.DB.Db.Create(&ad)
	return c.Status(200).JSON(ad)
}
