package service

import (
	"reflect"
	"testing"

	"github.com/ameykpatil/chatbot/model"
)

func TestSelectMostSuitableReply(t *testing.T) {
	tests := []struct {
		name       string
		replies    []model.Reply
		expReplies []model.Reply
	}{
		{
			name: "reply should be one of them",
			replies: []model.Reply{
				{
					Text:          "Hello",
					MinConfidence: 0.9,
					Intent:        "Greeting",
				},
				{
					Text:          "Hi",
					MinConfidence: 0.8,
					Intent:        "Greeting",
				},
			},
			expReplies: []model.Reply{
				{
					Text:          "Hello",
					MinConfidence: 0.9,
					Intent:        "Greeting",
				},
				{
					Text:          "Hi",
					MinConfidence: 0.8,
					Intent:        "Greeting",
				},
			},
		},
		{
			name:    "reply should be default",
			replies: []model.Reply{},
			expReplies: []model.Reply{
				{
					Text: "We couldn't understand it, can you please rephrase & try again?",
				},
			},
		},
	}

	for _, tt := range tests {
		reply := selectMostSuitableReply(tt.replies)

		found := false
		for _, expReply := range tt.expReplies {
			if reflect.DeepEqual(reply, expReply) {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("reply %v is not one of the expected replies", reply)
		}
	}

}
