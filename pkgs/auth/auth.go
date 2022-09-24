package auth

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	_ "golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	_ "golang.org/x/oauth2/google"
)

// GetGoogleChatOauthClient  get oauth client with desired scope privilege.
func GetGoogleChatOauthClient(ctx context.Context, serviceAccountKeyPath string) *http.Client {
	data, err := ioutil.ReadFile(serviceAccountKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	creds, err := google.CredentialsFromJSON(
		ctx,
		data,
		"https://www.googleapis.com/auth/chat.bot",
	)
	if err != nil {
		log.Fatal(err)
	}

	return oauth2.NewClient(ctx, creds.TokenSource)
}

// TODO: Check if user already voted. 
func hasVoted(user string) bool {
	return false
}