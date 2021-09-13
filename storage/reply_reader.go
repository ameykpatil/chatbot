package storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ameykpatil/chatbot/model"
)

const (
	database   = "chatbot"
	collection = "replies"
)

// ReplyReader struct to help read activities from storage
type ReplyReader struct {
	mongoClient *mongo.Client
}

// NewReplyReader create new instance of ReplyReader
func NewReplyReader(mongoClient *mongo.Client) *ReplyReader {
	return &ReplyReader{
		mongoClient: mongoClient,
	}
}

// GetByIntent returns replies for a given intent
func (rr *ReplyReader) GetByIntent(ctx context.Context, intent model.Intent) ([]model.Reply, error) {

	collection := rr.mongoClient.Database(database).Collection(collection)

	filter := bson.D{
		{"intent", intent.Name},
		{"min_confidence", bson.D{{"$lte", intent.Confidence}}},
	}

	var replies []model.Reply
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Printf("error while finding records from mongo : %s", err)
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var reply model.Reply
		err := cur.Decode(&reply)
		if err != nil {
			log.Printf("error while decoding record into reply : %s", err)
			return nil, err
		}
		replies = append(replies, reply)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return replies, nil
}
