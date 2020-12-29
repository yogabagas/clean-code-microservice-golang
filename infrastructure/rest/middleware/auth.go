package middleware

import (
	"my-github/clean-code-microservice-golang/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

func NewAuthentication(rest interfaces.Rest) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
