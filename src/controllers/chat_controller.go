package controllers

import (
	"firebase-chat/src/configs"
	"firebase-chat/src/dtos"
	"firebase-chat/src/models"
	"firebase-chat/src/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func createChatHandler(c *fiber.Ctx) error {
	chat := models.Chat{}

	if err := c.BodyParser(&chat); err != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed saving chat!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if validationErr := configs.Validator.Struct(&chat); validationErr != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed saving chat!",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := services.CreateChat(chat)
	if err != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed saving chat!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dtos.Response{
		Status:  http.StatusCreated,
		Message: "Successfully created a chat!",
		Data:    &fiber.Map{"data": result.ID},
	}
	return c.Status(http.StatusCreated).JSON(response)
}

func sendMessageHandler(c *fiber.Ctx) error {
	message := models.Message{}

	if err := c.BodyParser(&message); err != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed sending new message!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if validationErr := configs.Validator.Struct(&message); validationErr != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed sending new message!",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := services.SendMessage(message)
	if err != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed sending new message!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dtos.Response{
		Status:  http.StatusCreated,
		Message: "Successfully sent a message!",
		Data:    &fiber.Map{"data": result.ID},
	}
	return c.Status(http.StatusCreated).JSON(response)
}

func GetUserChatsHandler(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	result, err := services.GetUserChats(userID)
	if err != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed retrieving user's chats!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dtos.Response{
		Status:  http.StatusOK,
		Message: "Successfully retrieved user's chats!",
		Data:    &fiber.Map{"data": result},
	}
	return c.Status(http.StatusOK).JSON(response)
}

func AddChatController(router fiber.Router) {
	root := "/chats"

	router.Post(root, createChatHandler)
	router.Post(root+"/message", sendMessageHandler)
	router.Get(root+"/:user_id", GetUserChatsHandler)
}
