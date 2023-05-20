package models

type Chat struct {
	ID             string    `json:"id,omitempty" firestore:"id,omitempty"`
	ParticipantsID []string  `json:"participants_id" validate:"required" firestore:"participants_id"`
	Messages       []Message `json:"messages" firestore:"messages,omitempty"`
}
