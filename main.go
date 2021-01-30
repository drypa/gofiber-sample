package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"log"
)

func main() {

	app := fiber.New()

	app.Get("/a", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello")
	})

	authConfig := basicauth.Config{Authorizer: func(username string, pass string) bool {
		return false
	}}

	auth := app.Group("/auth")

	auth.Get("/a", func(ctx *fiber.Ctx) error {
		return ctx.SendString(" it's /auth/a")
	})

	secure := app.Group("/secure")

	secure.Get("/b", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello")
	})

	secure.Use(basicauth.New(authConfig))

	log.Print(app.Listen(":5000"))
}
