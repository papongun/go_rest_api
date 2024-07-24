package exception

import (
	"github.com/gofiber/fiber/v2"

	"github.com/papongun/go_todo/dto"
)

func HandleError(ctx *fiber.Ctx, err error) error {
	if validationError, ok := err.(ValidationError); ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{
			Error: validationError.parseErrorMessage(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{
		Error: "Internal Server Error",
	})
}
