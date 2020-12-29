package group

import (
	"my-github/clean-code-microservice-golang/infrastructure/rest/handler/v1/student"
	"my-github/clean-code-microservice-golang/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

func InitStudentV1(rest interfaces.Rest, v1 fiber.Router) {
	authGroup := v1.Group("/student")
	authGroup.Post("/", student.Create(rest))
}
