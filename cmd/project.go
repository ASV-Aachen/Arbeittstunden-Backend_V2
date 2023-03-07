package cmd

import (
	"github.com/ASV-Aachen/ArbeitsstundenDB/modules"
	db "github.com/ASV-Aachen/ArbeitsstundenDB/modules/DB"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllProject(c *fiber.Ctx, database *gorm.DB) error {
	var currentProjects []db.Project
	database.Model(&db.Project{}).Find(currentProjects)

	return c.Status(fiber.StatusOK).JSON(currentProjects)
}

func NewProject(c *fiber.Ctx, database *gorm.DB) error {
	json := new(modules.JsonProject)
	if err := c.BodyParser(json); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var currentSeason db.Season
	database.Model(&db.Season{Year: json.First_season})

	database.Create(&db.Project{
		Name:         json.Name,
		Description:  json.Description,
		First_season: currentSeason,
	})

	return c.Status(fiber.StatusOK).SendStatus(200)
}

func SinglePrjectGetInfos(c *fiber.Ctx, database *gorm.DB) error {
	searched_id := c.Params("projectID")

	var currentProjects db.Project
	database.Model(&db.Project{Id: searched_id}).First(currentProjects)

	return c.Status(fiber.StatusOK).JSON(currentProjects)
}

func UpdateProject(c *fiber.Ctx, database *gorm.DB) error {
	searched_id := c.Params("projectID")
	var currentProjects db.Project
	database.Model(&db.Project{Id: searched_id}).First(currentProjects)

	json := new(modules.JsonProject)
	if err := c.BodyParser(json); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var currentSeason db.Season
	database.Model(&db.Season{Year: json.First_season})

	currentProjects.Name = json.Name
	currentProjects.Description = json.Description
	currentProjects.First_season = currentSeason

	database.Save(currentProjects)

	return c.Status(fiber.StatusOK).SendStatus(200)
}
