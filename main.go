package main

import (
	"crypto/tls"
	"net/http"

	"github.com/ASV-Aachen/ArbeitsstundenDB/cmd"
	"github.com/ASV-Aachen/ArbeitsstundenDB/setup"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func main() {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	app := fiber.New()
	app.Use(cors.New())
	app.Use(
		logger.New(),
	)

	// Test handler
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("App running")
	})

	var db *gorm.DB = setup.SetUpMariaDB()
	setup.DB_Migrate(db)

	api := app.Group("/arbeitsstunden/V2/api")

	api.Use(
		setup.Check_IsUserLoggedIn,
	)

	member := api.Group("/member")
	member.Get("", func(c *fiber.Ctx) error { return cmd.GetAllMember(c, db) })
	member.Get("/WhoAmI", func(c *fiber.Ctx) error {

		return c.SendString("")
	})

	singleMember := member.Group("/:memberID")
	singleMember.Get("/", func(c *fiber.Ctx) error { return cmd.GetMember(c, db) })
	singleMember.Patch("/", func(c *fiber.Ctx) error { return cmd.MemberBearbeiten(c, db) })

	reduction := singleMember.Group("/reduction")
	reduction.Get("/", func(c *fiber.Ctx) error { return cmd.Reduction_getForSpecificUser(c, db) })
	reduction.Post("/", func(c *fiber.Ctx) error { return cmd.Reduction_Add(c, db) })
	reduction.Patch("/:reductionID", func(c *fiber.Ctx) error { return cmd.Reduction_Update(c, db) })

	projects := api.Group("/projects")
	projects.Get("", func(c *fiber.Ctx) error { return cmd.GetAllProject(c, db) })
	projects.Post("", func(c *fiber.Ctx) error { return cmd.NewProject(c, db) })

	singleProject := projects.Group("/:projectID")
	singleProject.Get("", func(c *fiber.Ctx) error { return cmd.SinglePrjectGetInfos(c, db) })
	singleProject.Patch("", func(c *fiber.Ctx) error { return cmd.UpdateProject(c, db) })

	arbeiten := projects.Group("/arbeiten")
	arbeiten.Get("", func(c *fiber.Ctx) error { return cmd.GetAllArbeitenForProject(c, db) })
	arbeiten.Post("", func(c *fiber.Ctx) error { return cmd.AddNewArbeitenForProject(c, db) })

	singleArbeiten := arbeiten.Group("/:arbeitsID")
	singleArbeiten.Post("", func(c *fiber.Ctx) error { return cmd.EinzelneArbeitbearbeiten(c, db) })
	singleArbeiten.Delete("", func(c *fiber.Ctx) error { return cmd.ArbeiteEntfernen(c, db) })

	arbeiten_takel := arbeiten.Group("/takel")
	arbeiten_takel.Get("", func(c *fiber.Ctx) error { return cmd.GetAllArbeitenNotChecked(c, db) })

	app.Listen(":3000")
}
