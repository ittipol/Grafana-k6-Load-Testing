package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func VersionV1(c *fiber.Ctx) error {
	c.Set("Version", "v1")
	return c.Next()
}

func Logger(c *fiber.Ctx) error {

	return c.Next()
}

func Auth(c *fiber.Ctx) error {

	return c.Next()
}
