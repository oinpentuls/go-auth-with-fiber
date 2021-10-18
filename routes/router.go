package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oinpentuls/auth-with-fiber/controller"
)

func Setup(app *fiber.App) {
	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)
	app.Post("/logout", controller.Logout)
	app.Get("/user", controller.User)
}
