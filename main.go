package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		// time.Sleep(time.Second * 1)
		return c.JSON("OK")
	})

	app.Listen(":5000")

}
