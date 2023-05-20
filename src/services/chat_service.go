package services

import (
	"firebase-chat/src/models"
	"firebase-chat/src/repositories"

	"cloud.google.com/go/firestore"
)

func CreateChat(chat models.Chat) (*firestore.DocumentRef, error) {
	return repositories.CreateChat(chat)
}

func SendMessage(message models.Message) (*firestore.DocumentRef, error) {
	return repositories.SendMessage(message)
}

func GetUserChats(userID string) ([]models.Chat, error) {
	return repositories.GetUserChats(userID)
}
