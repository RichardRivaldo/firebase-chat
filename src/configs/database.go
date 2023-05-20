package configs

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func CreateDBConnection() *firestore.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conf := &firebase.Config{
		ProjectID: GetEnv("FIREBASE_PROJECT_ID"),
	}

	opt := option.WithCredentialsFile(GetEnv("FIREBASE_CREDS"))

	app, err := firebase.NewApp(ctx, conf, opt)

	if err != nil {
		log.Fatalln("Error initializing Firebase, reason: ", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln("Error creating database, reason: ", err)
	}

	return client
}

func GetCollection(client *firestore.Client, path string) *firestore.CollectionRef {
	return client.Collection(path)
}

var DB = CreateDBConnection()
