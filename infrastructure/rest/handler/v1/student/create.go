package student

import (
	"my-github/clean-code-microservice-golang/domain/model"
	"my-github/clean-code-microservice-golang/internal/interfaces"
	"my-github/clean-code-microservice-golang/internal/response"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
)

func Create(rest interfaces.Rest) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var student *model.Student

		if err := c.BodyParser(&student); err != nil {
			return response.Fail(c, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, err)
		}

		if ok, err := isRequestValid(*student); !ok {
			return response.Fail(c, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, err)
		}
		appController := rest.GetAppController()
		err := appController.Student.CreateStudent(c.Context(), student)
		if err != nil {
			return response.Fail(c, http.StatusInternalServerError, http.StatusInternalServerError, err)
		}
		return response.Success(c, http.StatusOK, "success")
	}
}

func isRequestValid(req model.Student) (bool, error) {

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.StudentID, validation.Required),
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.DateOfBirth, validation.Required)); err != nil {
		return false, err
	}
	return true, nil
}
