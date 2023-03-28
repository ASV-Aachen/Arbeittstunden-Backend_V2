package cmd

import (
	"github.com/ASV-Aachen/Arbeittstunden-Backend_V2/modules"
	db "github.com/ASV-Aachen/Arbeittstunden-Backend_V2/modules/DB"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllArbeitenForProject(c *fiber.Ctx, database *gorm.DB) error {
	searched_id := c.Params("projectID")

	var currentProject db.Project
	database.Model(&db.Project{Id: searched_id}).First(&currentProject)

	var currentArbeiten db.Project_item
	database.Model(&db.Project_item{Project: currentProject}).Find(&currentArbeiten)

	return c.Status(fiber.StatusOK).JSON(currentArbeiten)
}

func AddNewArbeitenForProject(c *fiber.Ctx, database *gorm.DB) error {
	searched_id := c.Params("projectID")
	var currentProject db.Project
	database.Model(&db.Project{Id: searched_id}).First(&currentProject)

	json := new(modules.Json_Project_Item)
	if err := c.BodyParser(json); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var currentSeason db.Season
	database.Model(&db.Season{Year: json.Season})

	var temp db.Project_item
	temp.Project = currentProject
	temp.Season = currentSeason
	temp.Date = json.Date
	temp.Title = json.Title
	temp.Description = json.Description
	temp.Approved = json.Approved
	temp.Countable = json.Countable

	for _, i := range json.Work {
		var temp_work db.Project_item_hour
		temp_work.Duration = i.Duration

		var currentUsers db.User
		database.Model(&db.User{Id: i.Member}).First(&currentUsers)
		temp_work.Member = currentUsers
		temp_work.Project_item = temp
		database.Save(temp_work)
	}

	database.Create(temp)

	return c.Status(fiber.StatusOK).SendStatus(200)
}

func EinzelneArbeitbearbeiten(c *fiber.Ctx, database *gorm.DB) error {
	searched_id := c.Params("arbeitsID")
	var currentArbeit db.Project_item
	database.Model(&db.Project_item{Id: searched_id}).First(&currentArbeit)

	json := new(modules.Json_Project_Item)
	if err := c.BodyParser(json); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var currentSeason db.Season
	database.Model(&db.Season{Year: json.Season})

	currentArbeit.Season = currentSeason
	currentArbeit.Date = json.Date
	currentArbeit.Title = json.Title
	currentArbeit.Description = json.Description
	currentArbeit.Approved = json.Approved
	currentArbeit.Countable = json.Countable

	for _, i := range json.Work {
		var temp_work db.Project_item_hour
		temp_work.Duration = i.Duration

		var currentUsers db.User
		database.Model(&db.User{Id: i.Member}).First(&currentUsers)
		temp_work.Member = currentUsers

		temp_work.Project_item = currentArbeit
		database.Save(temp_work)
	}

	database.Save(currentArbeit)

	return c.Status(fiber.StatusOK).SendStatus(200)
}

func ArbeiteEntfernen(c *fiber.Ctx, database *gorm.DB) error {
	searched_id := c.Params("arbeitsID")
	var currentArbeit db.Project_item
	database.Model(&db.Project_item{Id: searched_id}).First(&currentArbeit)

	database.Delete(currentArbeit)

	return c.Status(fiber.StatusOK).SendStatus(200)
}

func GetAllArbeitenNotChecked(c *fiber.Ctx, database *gorm.DB) error {
	var currentArbeit []db.Project_item
	database.Model(&db.Project_item{Approved: false}).First(&currentArbeit)

	return c.Status(fiber.StatusOK).JSON(currentArbeit)
}
