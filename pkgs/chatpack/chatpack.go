package chatpack

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nurhun/google-vote-chatbot-golang/pkgs/auth"
	chat "google.golang.org/api/chat/v1"
	"google.golang.org/api/option"
)

var VoteCount int
var param = chat.ActionParameter{
	Key:   "count",
	Value: fmt.Sprint(VoteCount),
}

var message chat.Message
var msg_created string

// GetNewMSGService setup client to deal with chat.google.com.
func GetNewMSGService(ctx context.Context) *chat.SpacesMessagesService {
	chatService := getChatService(ctx)
	newMsgService := chat.NewSpacesMessagesService(chatService)
	return newMsgService
}

func getChatService(ctx context.Context) *chat.Service {
	oauthClient := auth.GetGoogleChatOauthClient(ctx, os.Getenv("SA_KEY_PATH"))

	chatService, err := chat.NewService(ctx, option.WithHTTPClient(oauthClient))
	if err != nil {
		log.Printf("chatService error: %v", err)
	}
	return chatService
}

// CreateMSG creates new messages.
func CreateMSG(title string, msgService *chat.SpacesMessagesService) error {
	VoteCount = 0
	msg := createChatCard(VoteCount, title)
	message, err := msgService.Create(os.Getenv("SPACE"), msg).Do()
	if err != nil {
		log.Printf("error create new message: %v\n", err)
		return err
	}

	msg_created = message.Name

	log.Printf("created messageeeeeeeee: %v\n", message.Name)

	return nil
}

// UpdateMSG updates existing message to increase voteCount.
func UpdateMSG(title string, msgService *chat.SpacesMessagesService) error {

	VoteCount += 1
	msg := updateChatCard(VoteCount, title)

	log.Printf("message to be updatedddddddddd: %v\n", msg_created)

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
		"cards": [
			{
				"header": 
					{
						"title": "%s"
					},
				"sections": [
					{
						"widgets": [
							{
								"textParagraph": {"text": "%d votes!"}
							},
							{
								"buttons": [
									{
										"textButton": {
											"text": "+1",
											"onClick": {
												"action": {
													"actionMethodName": "upvote",
													"parameters": [
														{
															"key": "count",
															"value": "%s"
														}
													]
												}
											}
										}
									},{
										"textButton": {
											"text": "NEW",
											"onClick": {
												"action": {
													"actionMethodName": "newvote"
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

func updateChatCard(vote_count int, title string) *chat.Message {

	param.Value = fmt.Sprint(vote_count)

	cards := `{
		"actionResponse": 
			{
				"type": "UPDATE_MESSAGE"
			},
		"cards": [
			{
				"header": 
					{
						"title": "%s"
					},
				"sections": [
					{
						"widgets": [
							{
								"textParagraph": {"text": "%d votes!"}
							},
							{
								"buttons": [
									{
										"textButton": {
											"text": "+1",
											"onClick": {
												"action": {
													"actionMethodName": "upvote",
													"parameters": [
														{
															"key": "count",
															"value": "%s"
														}
													]
												}
											}
										}
									},{
										"textButton": {
											"text": "NEW",
											"onClick": {
												"action": {
													"actionMethodName": "newvote"
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
