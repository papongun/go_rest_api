package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/papongun/go_todo/config"
	auth_c "github.com/papongun/go_todo/controller/auth"
	"github.com/papongun/go_todo/repository"
	auth_s "github.com/papongun/go_todo/service/auth"
)

func main() {
	app := fiber.New()

	db := config.InitDatabase()
	userRepo := repository.NewUserRepository(db)
	authRegService := auth_s.NewAuthRegisterService(userRepo)
	authRegController := auth_c.NewUserRegisterContoller(authRegService)

	v1 := app.Group("/v1")

	authHandlers := v1.Group("/auth")
	authHandlers.Post("/login", authRegController.Register)

	app.Listen(":8080")
}
