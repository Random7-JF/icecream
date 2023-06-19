package server

import (
	"github.com/Random-7/icecream/app/config"
	"github.com/gofiber/fiber/v2"
)

func Serve(config *config.App) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(config.IP + ":" + config.Port)
}
