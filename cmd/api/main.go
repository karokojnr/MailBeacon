package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"MailBeacon/internal/database"
	"MailBeacon/internal/mailer"
	"MailBeacon/internal/newsletter"
	"MailBeacon/internal/pubsub"
	transportHTTP "MailBeacon/internal/transport/http"

	googlePubSub "cloud.google.com/go/pubsub"

	_ "github.com/joho/godotenv/autoload"
)

const version = ""

var (
	projectId = os.Getenv("PROJECT_ID")
)

func Run() error {

	// Create a new database connection
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	// Create a new pubsub client
	client, err := googlePubSub.NewClient(context.Background(), projectId)
	if err != nil {
		log.Fatalf("Failed to create pubsub client: %v", err)
	}
	defer client.Close()
	pSub := pubsub.NewGooglePubSub(client)

	// Create a new mailer client
	sendgridClient := mailer.NewSendGrid()

	// Create a new newsletter service
	newsletterService := newsletter.NewNewsletterService(db, pSub, sendgridClient)

	// Create a new http handler
	httpHandler := transportHTTP.NewHandler(newsletterService)
	log.Printf("Server is running on: %v, version: %v", httpHandler.Server.Addr, version)
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
