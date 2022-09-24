package chatpack

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nurhun/google-vote-chatbot-golang/internal/auth"
	chat "google.golang.org/api/chat/v1"
	"google.golang.org/api/option"
)

var voteCount int
var param = chat.ActionParameter{
	Key:   "count",
	Value: fmt.Sprint(voteCount),
}

var message chat.Message
var msg_created string

// GetNewMSGService setup client to deal with chat.google.com.
func GetNewMSGService(ctx context.Context) *chat.SpacesMessagesService {
	chatService, _ := getChatService(ctx)
	newMsgService := chat.NewSpacesMessagesService(chatService)
	return newMsgService
}

func getChatService(ctx context.Context) (*chat.Service, error) {
	oauthClient := auth.GetGoogleChatOauthClient(ctx, os.Getenv("SA_KEY_PATH"))

	chatService, err := chat.NewService(ctx, option.WithHTTPClient(oauthClient))
	if err != nil {
		log.Printf("chatService error: %v", err)
		return nil, err
	}
	return chatService, nil
}

// CreateMSG creates new messages.
func CreateMSG(title string, space string, msgService *chat.SpacesMessagesService) error {
	voteCount = 0
	msg := createChatCard(voteCount, title)
	message, err := msgService.Create(space, msg).Do()
	if err != nil {
		log.Printf("error create new message: %v\n", err)
		return err
	}

	msg_created = message.Name

	log.Printf("created message: %v\n", message.Name)

	return nil
}

// UpdateMSG updates existing message to increase voteCount.
func UpdateMSG(title string, msgService *chat.SpacesMessagesService) error {

	voteCount += 1
	msg := updateChatCard(voteCount)

	message, err := msgService.Update(msg_created, msg).UpdateMask("cards").Do()
	if err != nil {
		log.Printf("error update message: %v", err)
		return err
	}

	log.Printf("message updated: %v\n", message.Name)

	return nil
}

func createChatCard(vote_count int, title string) *chat.Message {

	param.Value = fmt.Sprint(vote_count)

	cards := `{
		"cards":[
		   {
			  "header":{
				 "title":"%s"
			  },
			  "sections":[
				 {
					"widgets":[
					   {
						  "textParagraph":{
							 "text":"%d votes!"
						  }
					   },
					   {
						  "buttons":[
							 {
								"textButton":{
								   "text":"+1",
								   "onClick":{
									  "action":{
										 "actionMethodName":"upvote",
										 "parameters":[
											{
											   "key":"count",
											   "value":"%s"
											}
										 ]
									  }
								   }
								}
							 },
							 {
								"textButton":{
								   "text":"NEW",
								   "onClick":{
									  "action":{
										 "actionMethodName":"newvote"
									  }
								   }
								}
							 }
						  ]
					   }
					]
				 }
			  ]
		   }
		]
	 }`

	outputString := fmt.Sprintf(cards, title, vote_count, param.Value)
	json.Unmarshal([]byte(outputString), &message)

	return &message
}

func updateChatCard(vote_count int) *chat.Message {

	param.Value = fmt.Sprint(vote_count)

	cards := `{
		"actionResponse":{
		   "type":"UPDATE_MESSAGE"
		},
		"cards":[
		   {
			  "sections":[
				 {
					"widgets":[
					   {
						  "textParagraph":{
							 "text":"%d votes!"
						  }
					   },
					   {
						  "buttons":[
							 {
								"textButton":{
								   "text":"+1",
								   "onClick":{
									  "action":{
										 "actionMethodName":"upvote",
										 "parameters":[
											{
											   "key":"count",
											   "value":"%s"
											}
										 ]
									  }
								   }
								}
							 },
							 {
								"textButton":{
								   "text":"NEW",
								   "onClick":{
									  "action":{
										 "actionMethodName":"newvote"
									  }
								   }
								}
							 }
						  ]
					   }
					]
				 }
			  ]
		   }
		]
	 }`

	outputString := fmt.Sprintf(cards, vote_count, param.Value)
	json.Unmarshal([]byte(outputString), &message)

	return &message
}
