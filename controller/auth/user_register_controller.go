package auth

import (
	"github.com/gofiber/fiber/v2"

	generic "github.com/papongun/go_todo/dto"
	dto "github.com/papongun/go_todo/dto/auth"
	"github.com/papongun/go_todo/exception"
	service "github.com/papongun/go_todo/service/auth"
)

type UserRegisterContoller struct {
	s service.AuthRegisterService
}

func NewUserRegisterContoller(s service.AuthRegisterService) *UserRegisterContoller {
	return &UserRegisterContoller{s: s}
}

func (c *UserRegisterContoller) Register(ctx *fiber.Ctx) error {
	var registerRequest dto.UserRegisterRequest
	if err := ctx.BodyParser(&registerRequest); err != nil {
		return exception.HandleError(ctx, err)
	}

	response, err := c.s.Register(&registerRequest)
	if err != nil {
		return exception.HandleError(ctx, err)
	}
	return ctx.Status(fiber.StatusCreated).JSON(generic.Response[dto.UserRegisterResponse]{Message: "Register success", Data: *response})
}
