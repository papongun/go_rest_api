package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/papongun/go_todo/config"
	"github.com/papongun/go_todo/router"
)

func main() {
	app := fiber.New()

	db := config.InitDatabase()
	v1 := app.Group("/v1")

	router.InitAuthRoute(v1, db)

	app.Listen(":8080")
}
