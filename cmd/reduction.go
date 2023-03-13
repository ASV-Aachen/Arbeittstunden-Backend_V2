package cmd

import (
	"github.com/ASV-Aachen/Arbeittstunden-Backend_V2/modules"
	db "github.com/ASV-Aachen/Arbeittstunden-Backend_V2/modules/DB"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Reduction_getForSpecificUser(c *fiber.Ctx, database *gorm.DB) error {
	searched_id := c.Params("memberID")
	var currentUsers db.User
	database.Model(&db.User{Id: searched_id}).Find(currentUsers)

	var currentReductions []db.Reduction
	database.Model(&db.Reduction{Member: currentUsers}).Find(currentReductions)

	return c.Status(fiber.StatusOK).JSON(currentReductions)
}

func Reduction_Add(c *fiber.Ctx, database *gorm.DB) error {
	searched_id := c.Params("memberID")
	var currentUsers db.User
	database.Model(&db.User{Id: searched_id}).Find(currentUsers)

	json := new(modules.JsonReduction)
	if err := c.BodyParser(json); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var currentSeason db.Season
	database.Model(&db.Season{Year: json.Season}).First(currentSeason)

	database.Create(&db.Reduction{
		Season:               currentSeason,
		Member:               currentUsers,
		Reduction_in_percent: json.Reduction_in_percent,
		Note:                 json.Note,
	})

	return c.Status(fiber.StatusOK).SendStatus(200)
}

func Reduction_Update(c *fiber.Ctx, database *gorm.DB) error {
	searched_id_reduction := c.Params("reductionID")

	var currentRedutcion db.Reduction
	database.Model(&db.Reduction{ID: searched_id_reduction}).First(currentRedutcion)

	json := new(modules.JsonReduction)
	if err := c.BodyParser(json); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var currentSeason db.Season
	database.Model(&db.Season{Year: json.Season}).First(currentSeason)

	currentRedutcion.Season = currentSeason
	currentRedutcion.Reduction_in_percent = json.Reduction_in_percent
	currentRedutcion.Note = json.Note

	database.Save(currentRedutcion)

	return c.SendString("DO MAGIC")
}
