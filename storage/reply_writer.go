package storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ameykpatil/chatbot/model"
)

// DefaultReplies holds some replies to load in db
var DefaultReplies = []model.Reply{
	{
		Text:          "Hello",
		MinConfidence: 0.9,
		Intent:        "Greeting",
	},
	{
		Text:          "How can I help you?",
		MinConfidence: 0.8,
		Intent:        "Greeting",
	},
	{
		Text:          "I can connect you with our customer executive, would that be fine?",
		MinConfidence: 0.7,
		Intent:        "I want to speak with a human",
	},
	{
		Text:          "Great",
		MinConfidence: 0.6,
		Intent:        "Affirmative",
	},
	{
		Text:          "I am sorry that you are facing issues, can you please restart the app?",
		MinConfidence: 0.4,
		Intent:        "Login problems",
	},
	{
		Text:          "I am glad I was able to help",
		MinConfidence: 0.6,
		Intent:        "Thank you",
	},
	{
		Text:          "Have a great rest of the day",
		MinConfidence: 0.5,
		Intent:        "Goodbye",
	},
	{
		Text:          "Ciao",
		MinConfidence: 0.9,
		Intent:        "Goodbye",
	},
}

// LoadRepliesToDB loads provided replies to database
// This is just a workaround for loading initial data
// It can be structured in a more proper way similar to ReplyReader
func LoadRepliesToDB(ctx context.Context, mongoClient *mongo.Client, replies []model.Reply) error {

	collection := mongoClient.Database(database).Collection(collection)

	docs := make([]interface{}, 0)
	for _, v := range replies {
		v.ID = primitive.NewObjectID().String()
		docs = append(docs, v)
	}

	res, err := collection.InsertMany(ctx, docs)
	if err != nil {
		return err
	}

	log.Printf("loaded %d replies in db", len(res.InsertedIDs))
	return nil
}
