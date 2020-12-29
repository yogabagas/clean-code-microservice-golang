package response

import (
	"fmt"
	"math/rand"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
)

type (
	Causer interface {
		Cause() error
	}

	Response struct {
		RequestId string      `json:"request_id"`
		Content   interface{} `json:"content"`
		Error     *Error      `json:"error"`
		Status    int         `json:"status"`
	}

	Error struct {
		Code    int               `json:"code"`
		Message string            `json:"message"`
		Reasons validation.Errors `json:"reasons"`
	}
)

func (err *Error) Error() string {
	return fmt.Sprintf("error with code: %d; message: %s", err.Code, err.Message)
}

func Success(ctx *fiber.Ctx, status int, content interface{}) error {
	rand.Seed(time.Now().UTC().UnixNano())
	requestID := rand.Int63()
	ctx.Set("X-Request-ID", string(requestID))
	return ctx.JSON(&Response{
		RequestId: fmt.Sprintf("%d", requestID),
		Content:   content,
		Status:    status,
	})
}

func Fail(ctx *fiber.Ctx, status, errCode int, err error) error {
	var (
		message = err.Error()
		reason  = validation.Errors{}
	)

	rand.Seed(time.Now().UTC().UnixNano())
	requestID := rand.Int63()

	if cause, ok := err.(Causer); ok {
		err = cause.Cause()
	}

	if valid, ok := err.(validation.Errors); ok {
		message = "there`s some validation issues in request attributes"
		reason = valid
	}
	ctx.Set("X-Request-ID", string(requestID))
	ctx.Status(status)

	return ctx.JSON(&Response{
		RequestId: fmt.Sprintf("%d", requestID),
		Status:    status,
		Error: &Error{
			Code:    errCode,
			Message: message,
			Reasons: reason,
		},
	})

}
