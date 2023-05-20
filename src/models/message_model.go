package models

type Message struct {
	ID       string `json:"id,omitempty" firestore:"id,omitempty"`
	ChatID   string `json:"chat_id,omitempty" validate:"required" firestore:"chat_id,omitempty"`
	SenderID string `json:"sender_id" validate:"required" firestore:"sender_id,omitempty"`
	SentAt   string `json:"sent_at" validate:"required" firestore:"sent_at,omitempty"`
	Content  string `json:"content" validate:"required" firestore:"content,omitempty"`
}
