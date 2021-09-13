package service

import (
	"context"
	"math/rand"

	"github.com/ameykpatil/chatbot/client"
	"github.com/ameykpatil/chatbot/model"
	"github.com/ameykpatil/chatbot/storage"
)

// ReplyService holds properties required for serving reply
type ReplyService struct {
	intentClient client.IntentClient
	replyReader  storage.ReplyReader
}

// NewReplyService creates an instance of ReplyService
func NewReplyService(rr storage.ReplyReader, ic client.IntentClient) *ReplyService {
	return &ReplyService{
		intentClient: ic,
		replyReader:  rr,
	}
}

// FindReply finds a reply for provided bot-id & message with the help of intent-api & replies in the db
func (rs *ReplyService) FindReply(ctx context.Context, botID, message string) (*model.Reply, error) {
	intentResponse, err := rs.intentClient.GetIntents(ctx, botID, message)
	if err != nil {
		return nil, err
	}

	// write logic here to select intent or intents to consider
	// can be taken out as a separate function
	intent := intentResponse.Intents[0]

	replies, err := rs.replyReader.GetByIntent(ctx, intent)
	if err != nil {
		return nil, err
	}

	// logic to select most suitable reply
	reply := selectMostSuitableReply(replies)

	return &reply, nil
}

// selectMostSuitableReply holds a logic of returning most suitable reply from given replies
// right now it returns it on random basis
func selectMostSuitableReply(replies []model.Reply) model.Reply {
	if len(replies) == 0 {
		return model.Reply{
			Text: "We couldn't understand it, can you please rephrase & try again?",
		}
	}

	randomIndex := rand.Intn(len(replies))
	return replies[randomIndex]
}
