package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func GetCors() func(*fiber.Ctx) error {
	config := cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}

	return cors.New(config)
}

func GetRecover() func(*fiber.Ctx) error {
	config := recover.Config{EnableStackTrace: true}
	return recover.New(config)
}
