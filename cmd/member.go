package cmd

import (
	"github.com/ASV-Aachen/Arbeittstunden-Backend_V2/modules"
	db "github.com/ASV-Aachen/Arbeittstunden-Backend_V2/modules/DB"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllMember(c *fiber.Ctx, database *gorm.DB) error {
	var currentUsers []db.User
	database.Model(&db.User{}).Find(&currentUsers)

	return c.Status(fiber.StatusOK).JSON(currentUsers)
}

func GetMember(c *fiber.Ctx, database *gorm.DB) error {
	searched_id := c.Params("memberID")
	var currentUsers db.User
	database.Model(&db.User{Id: searched_id}).Find(&currentUsers)

	return c.Status(fiber.StatusOK).JSON(currentUsers)
}

func MemberBearbeiten(c *fiber.Ctx, database *gorm.DB) error {
	searched_id := c.Params("memberID")
	json := new(modules.JsonMemberRolle)
	if err := c.BodyParser(json); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var currentUsers db.User
	database.Model(&db.User{Id: searched_id}).Find(&currentUsers)
	currentUsers.Rolle = json.Rolle
	database.Save(currentUsers)

	return c.SendString("DO MAGIC")
}
