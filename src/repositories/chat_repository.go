package repositories

import (
	"context"
	"firebase-chat/src/configs"
	"firebase-chat/src/models"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

var chatRef = configs.GetCollection(configs.DB, "chats")

func CreateChat(chat models.Chat) (*firestore.DocumentRef, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ref := chatRef.NewDoc()
	chat.ID = ref.ID

	_, err := ref.Set(ctx, chat)
	return ref, err
}

func SendMessage(message models.Message) (*firestore.DocumentRef, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ref := chatRef.Doc(message.ChatID).Collection("messages").NewDoc()
	message.ID = ref.ID

	_, err := ref.Set(ctx, message)
	return ref, err
}

func GetChatMessages(chatID string) ([]models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var messages []models.Message

	iter := configs.GetCollection(configs.DB, "chats/"+chatID+"/messages").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		var message models.Message
		if err := doc.DataTo(&message); err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func GetUserChats(userID string) ([]models.Chat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var chats []models.Chat

	query := chatRef.Where("participants_id", "array-contains-any", []string{userID})
	iter := query.Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		var chat models.Chat
		if err := doc.DataTo(&chat); err != nil {
			return nil, err
		}

		messages, err := GetChatMessages(chat.ID)
		if err != nil {
			return nil, err
		}

		chat.Messages = messages
		chats = append(chats, chat)
	}
	return chats, nil
}
