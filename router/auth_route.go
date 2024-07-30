package router

import (
	"github.com/gofiber/fiber/v2"

	auth_c "github.com/papongun/go_todo/controller/auth"
)

func InitAuthRoute(router fiber.Router) {
	authRegController := auth_c.GetUserRegisterContoller()
	authRoute := router.Group("/auth")
	authRoute.Post("/login", authRegController.Register)
}
