package group

import (
	"my-github/clean-code-microservice-golang/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

func InitV1(rest interfaces.Rest, root fiber.Router) fiber.Router {
	return root.Group("/v1")
}
