package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nurhun/google-vote-chatbot-golang/pkgs/chatpack"

	chat "google.golang.org/api/chat/v1"
)

var Event chat.DeprecatedEvent

func main() {
	log.Println("Hello World!")

	ctx := context.Background()
	newMsgService := chatpack.GetNewMSGService(ctx)

	handler := func(w http.ResponseWriter, r *http.Request) {
		// Checking Request method.
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Parsing request body.
		if err := json.NewDecoder(r.Body).Decode(&Event); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		switch Event.Type {
		case "ADDED_TO_SPACE":
			if Event.Space.Type != "ROOM" {
				break
			}
			fmt.Fprint(w, `{"text":"thanks for adding me!"}`)

		case "MESSAGE":
			if Event.Message.SlashCommand != nil {
				switch Event.Message.SlashCommand.CommandId {
				case 1:
					chatpack.CreateMSG(fmt.Sprintf("Vote on last poll by: %s", Event.User.DisplayName), newMsgService)
					// chatpack.CreateMSG(fmt.Sprintf("Vote on last poll by: %s", "nur"), newMsgService)
				default:
					fmt.Fprintf(w, `{"text":"No such option!"}`)
				}
			} else {
				fmt.Fprintf(w, `{"text":"No reply for %v"}`, Event.Message.Text)
			}

		case "CARD_CLICKED":
			log.Printf("Event: %v", Event.Action.ActionMethodName)

			if Event.Action.ActionMethodName == "newvote" {
				chatpack.CreateMSG(fmt.Sprintf("Vote on last poll by: %s", Event.User.DisplayName), newMsgService)
				// chatpack.CreateMSG(fmt.Sprintf("Vote on last poll by: %s", "nur"), newMsgService)
			} else if Event.Action.ActionMethodName == "upvote" {
				chatpack.UpdateMSG(fmt.Sprintf("Vote on last poll by: %s", Event.User.DisplayName), newMsgService)
				// chatpack.UpdateMSG(fmt.Sprintf("Vote on last poll by: %s", "nur"), newMsgService)
			}

			// TODO: hasVoted func.
			// TODO: isAuthorizerd function to check if in certain group.
		}
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), http.HandlerFunc(handler)))
}
