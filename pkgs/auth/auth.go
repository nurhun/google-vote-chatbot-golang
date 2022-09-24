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

func hasVoted(user string) bool {
	return false
}

// func getGoogleChatOauthClient() oauth2.TokenSource {
// 	ctx := context.Background()

// 	// Your credentials should be obtained from the Google
// 	// Developer Console (https://console.developers.google.com).
// 	conf := &jwt.Config{
// 		Email: "xxx@developer.gserviceaccount.com",
// 		// The contents of your RSA private key or your PEM file
// 		// that contains a private key.
// 		// If you have a p12 file instead, you
// 		// can use `openssl` to export the private key into a pem file.
// 		//
// 		//    $ openssl pkcs12 -in key.p12 -passin pass:notasecret -out key.pem -nodes
// 		//
// 		// The field only supports PEM containers with no passphrase.
// 		// The openssl command will convert p12 keys to passphrase-less PEM containers.
// 		PrivateKey: []byte("-----BEGIN RSA PRIVATE KEY-----..."),
// 		Scopes: []string{
// 			"https://www.googleapis.com/auth/bigquery",
// 			"https://www.googleapis.com/auth/blogger",
// 		},
// 		TokenURL: google.JWTTokenURL,
// 		// If you would like to impersonate a user, you can
// 		// create a transport with a subject. The following GET
// 		// request will be made on the behalf of user@example.com.
// 		// Optional.
// 		Subject: "user@example.com",
// 	}
// 	// Initiate an http.Client, the following GET request will be
// 	// authorized and authenticated on the behalf of user@example.com.
// 	return conf.TokenSource(ctx)
// }

// func getGoogleChatOauthClient(serviceAccountKeyPath string) *http.Client {
// 	ctx := context.Background()
// 	data, err := ioutil.ReadFile(serviceAccountKeyPath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	creds, err := google.CredentialsFromJSON(
// 		ctx,
// 		data,
// 		"https://www.googleapis.com/auth/chat.bot",
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return oauth2.NewClient(ctx, creds.TokenSource)
// }
