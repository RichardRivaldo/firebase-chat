# firebase-chat

Backend Application for Chat with Firebase

# Description

Simple chat application made with Firebase and Go Fiber.

# Setup

1. Add these environment variables to `.env` file.

```
SERVER_HOST=<PORT>
SERVER_PORT=<HOST>

FIREBASE_CREDS=<JSON_CREDS_PATH>
FIREBASE_DB_URI=<FIRESTORE_PROJECT_ID>
```

2. Use `go run .` or `go build` to install all dependencies and run the application.

# Flow

1. The user will initialize a chat first. The application tracks all participants in the chatroom. This way, aside from one-to-one or personal communication, a one-to-all or group chat communication can be done.
2. If a user send a new message, the system will track the chat in which the user send the message in. This way, the recepients of the chat can be correctly found by the system.
3. The system can then list out all chats a user has, based on the list of the participants in every chats, along with the messages in those chats.

# A Bit About Firebase

First time using `Firebase`! Decided to use `Firestore` on top of `Realtime Database` because it has the ability to structure the database into collections. I use Firestore ability to store a `subcollection` (collection inside collection). In this case, each chat will have a list of messages.

# API Documentation

`root` = `/api/v1`

## Starting a New Chat

```
POST /chats
{
    "participants_id": [<user_id1>, <user_id1>, ...]
}
```

Returns `chat_id` for the newly created chat.

## Sending a New Message

```
POST /chats/message
{
    "chat_id": <chat_id>,
    "sender_id": <sender_id>,
    "sent_at": <iso_timestamp>,
    "content": <content_text>
}
```

Returns `message_id` for the newly created message.

## Get All Chats for User

```
GET /chats/:user_id
```

Returns a list of chats with corresponding messages in each chat.

# Things to Improve

1. Better and more scalable models design(?)
2. A bit more on the clarity of response returned by the endpoints.
