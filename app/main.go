package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/papongun/go_todo/router"
)

func main() {
	app := fiber.New()
	v1 := app.Group("/v1")
	router.InitAuthRoute(v1)

	app.Listen(":8080")
}
