package group

import (
	"my-github/clean-code-microservice-golang/infrastructure/rest/middleware"
	"my-github/clean-code-microservice-golang/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

func InitRoot(rest interfaces.Rest) fiber.Router {
	router := rest.GetRouter()
	return router.Group("/", middleware.NewAuthentication(rest))
}
