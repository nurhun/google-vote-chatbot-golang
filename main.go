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

var event chat.DeprecatedEvent

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
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		switch event.Type {
		case "ADDED_TO_SPACE":
			if event.Space.Type != "ROOM" {
				break
			}
			fmt.Fprint(w, `{"text":"thanks for adding me!"}`)

		case "MESSAGE":
			if event.Message.SlashCommand != nil {
				switch event.Message.SlashCommand.CommandId {
				case 1:
					chatpack.CreateMSG(fmt.Sprintf("Vote on last poll by: %s", event.User.DisplayName), newMsgService)
				default:
					fmt.Fprintf(w, `{"text":"No such option!"}`)
				}
			} else {
				fmt.Fprintf(w, `{"text":"No reply for %v"}`, event.Message.Text)
			}

		case "CARD_CLICKED":
			log.Printf("event: %v", event.Action.ActionMethodName)

			if event.Action.ActionMethodName == "newvote" {
				chatpack.CreateMSG(fmt.Sprintf("Vote on last poll by: %s", event.User.DisplayName), newMsgService)
			} else if event.Action.ActionMethodName == "upvote" {
				chatpack.UpdateMSG(fmt.Sprintf("Vote on last poll by: %s", event.User.DisplayName), newMsgService)
			}

			// TODO: hasVoted func.
			// TODO: isAuthorizerd function to check if in certain group.
		}
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), http.HandlerFunc(handler)))
}
