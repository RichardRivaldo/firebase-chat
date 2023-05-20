package routers

import (
	"firebase-chat/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetRouters(app *fiber.App) {
	root := app.Group("api/v1")

	controllers.AddChatController(root)
}
