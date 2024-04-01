package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryankert01/goapitest/database"
	"github.com/ryankert01/goapitest/models"

	"fmt"
	"reflect"
	"strconv"
)

func GetAds(c *fiber.Ctx) error {
	q := c.Queries()
	fmt.Println(q)
	var err error
	var offset, limit, age int
	var gender, country, platform string

	age, err = strconv.Atoi(q["age"])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	offset, err = strconv.Atoi(q["offset"])
	if err != nil {
		offset = 0
	}

	limit, err = strconv.Atoi(q["limit"])
	if err != nil {
		limit = 5
	}

	gender, country, platform = q["gender"], q["country"], q["platform"]
	var cond []models.Conditions
	database.DB.Db.Model(models.Conditions{}).Where("age_start <= ? AND age_end >= ?", age, age).Where("? = ANY(gender)", gender).Where("? = ANY(country)", country).Where("? = ANY(platform)", platform).Offset(offset).Limit(limit).Find(&cond)

	return c.Status(200).JSON(cond)
}

func CreateAd(c *fiber.Ctx) error {
	ad := new(models.AD)
	if err := c.BodyParser(ad); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if IsAdAlive(*ad) {
		ad.IsActive = true
	} else {
		ad.IsActive = false
	}

	fmt.Print(ad.Conditions[0])
	fmt.Println(ad.Conditions[0].Gender, reflect.TypeOf(ad.Conditions[0].Gender))

	database.DB.Db.Model(models.AD{}).Create(&ad)
	return c.Status(200).JSON(ad)
}
