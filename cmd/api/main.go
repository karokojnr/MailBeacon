package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"MailBeacon/internal/database"
	"MailBeacon/internal/newsletter"
	"MailBeacon/internal/pubsub"
	transportHTTP "MailBeacon/internal/transport/http"

	googlePubSub "cloud.google.com/go/pubsub"

	_ "github.com/joho/godotenv/autoload"
)

var (
	projectId = os.Getenv("GOOGLE_PROJECT_ID")
)

func Run() error {

	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	client, err := googlePubSub.NewClient(context.Background(), projectId)
	if err != nil {
		log.Fatalf("Failed to create pubsub client: %v", err)
	}
	defer client.Close()

	pSub := pubsub.NewGooglePubSub(client)

	newsletterService := newsletter.NewNewsletterService(db, pSub)

	httpHandler := transportHTTP.NewHandler(newsletterService)

	log.Printf("Server is running on: %v", httpHandler.Server.Addr)
	if err := httpHandler.Serve(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
