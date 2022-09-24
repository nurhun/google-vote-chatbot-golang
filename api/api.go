package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nurhun/google-vote-chatbot-golang/internal/chatpack"
	"google.golang.org/api/chat/v1"
)

// Handler hhtp handler.s
func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	newMsgService := chatpack.GetNewMSGService(ctx)

	// Checking Request method.
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Receive coming event
	var event chat.DeprecatedEvent

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
				chatpack.CreateMSG(fmt.Sprintf("Vote on last poll by: %s", event.User.DisplayName), event.Space.Name, newMsgService)
			default:
				fmt.Fprintf(w, `{"text":"No such option!"}`)
			}
		} else {
			chatpack.CreateMSG(fmt.Sprintf("%s: %s", event.User.DisplayName, event.Message.Text), event.Space.Name, newMsgService)
		}

	case "CARD_CLICKED":
		log.Printf("event: %v", event.Action.ActionMethodName)

		if event.Action.ActionMethodName == "newvote" {
			chatpack.CreateMSG(fmt.Sprintf("Vote on last poll by: %s", event.User.DisplayName), event.Space.Name, newMsgService)
		} else if event.Action.ActionMethodName == "upvote" {
			chatpack.UpdateMSG(fmt.Sprintf("Vote on last poll by: %s", event.User.DisplayName), newMsgService)
		}
	default:
		fmt.Fprintf(w, `{"text":"No such option!"}`)

		// TODO: hasVoted func.
		// TODO: isAuthorizerd func.
	}
}
