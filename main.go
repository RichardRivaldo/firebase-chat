package main

import (
	"firebase-chat/src/configs"
	"firebase-chat/src/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routers.SetRouters(app)

	app.Listen(configs.GetEnv("SERVER_HOST") + ":" + configs.GetEnv("SERVER_PORT"))
}
