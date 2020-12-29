package interfaces

import (
	"my-github/clean-code-microservice-golang/adapter/controller"

	"github.com/gofiber/fiber/v2"
)

type Rest interface {
	GetRouter() *fiber.App
	GetAppController() *controller.AppController
}
