package handlers

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/ryankert01/goapitest/database"
// 	"github.com/ryankert01/goapitest/models"
// )

// func Home(c *fiber.Ctx) error {
// 	return c.SendString("Hello, World!!!! Ryan!!! or bobo..")
// }

// func CreateFact(c *fiber.Ctx) error {
// 	fact := new(models.Fact)
// 	if err := c.BodyParser(fact); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}
// 	database.DB.Db.Create(&fact)
// 	return c.Status(200).JSON(fact)
// }

// func ListFacts(c *fiber.Ctx) error {
// 	var facts []models.Fact
// 	database.DB.Db.Find(&facts)
// 	return c.Status(200).JSON(facts)
// }
