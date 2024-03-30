package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryankert01/goapitest/database"
	"github.com/ryankert01/goapitest/models"
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

	// cond := datatypes.JSONQuery(ad.Conditions)

	// fmt.Printf("ad: %v\n", ad.Conditions)

	database.DB.Db.Create(&ad)
	return c.Status(200).JSON(ad)
}
