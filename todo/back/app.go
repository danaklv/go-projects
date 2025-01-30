package app

import (
	"log"
	"os"
	"td/back/models"
	repositories "td/back/repostitories"
	"td/back/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func Run() {
	err := repositories.ConnectToDb()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var engine = html.New("./front/templates/", ".html")

	var App = fiber.New(fiber.Config{
		Views: engine,
	})

	App.Static("/static", "./front/static")

	App.Get("/", func(c *fiber.Ctx) error {

		return c.Render("intra", nil)

	})
	App.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", nil)
	})
	App.Get("/register", func(c *fiber.Ctx) error {

		return c.Render("register", nil)
	})

	App.Post("/register", func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		name := c.FormValue("name")

		user := models.User{
			Username: name,
			Email:    email,
			Password: password,
		}

		services.RegisterUserService(user)

		return c.Redirect("/")

	})

	log.Fatal(App.Listen(":8080"))
}
