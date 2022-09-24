package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nurhun/google-vote-chatbot-golang/api"
)

func main() {
	log.Println("Hello World!")

	// Simple http router for single handler func.
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), http.HandlerFunc(api.Handler)))
}