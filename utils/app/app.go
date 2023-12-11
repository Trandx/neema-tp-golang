package app

import (
	"github.com/gofiber/fiber/v2"
	. "neema.co.za/rest/utils/app/middleware"
)

// type RouterCreator func(prefix string, handlers ...fiber.Handler) fiber.Router
// type Router fiber.Router
type App struct {
	*fiber.App
}

var app *App

func init() {
	app = &App{NewFiberApp()}
}

func NewFiberApp() *fiber.App {
	return fiber.New(fiber.Config{
		DisableKeepalive: true,
	})
}

//func NewRouter() RouterCreator {
//	return app.Group
//}

func Initialise() *App {
	app.Use(GetRecover())
	app.Use(GetCors())
	return app
}
