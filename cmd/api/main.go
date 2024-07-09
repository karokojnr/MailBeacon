package main

import (
	"fmt"

	"MailBeacon/internal/database"
	"MailBeacon/internal/newsletter"
	transportHTTP "MailBeacon/internal/transport/http"
)

func Run() error {
	fmt.Println("Starting server...")

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
	}

	// todo: add more services here and add db to them
	newsletterService := newsletter.NewNewsletterService(db)

	httpHandler := transportHTTP.NewHandler(newsletterService)
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
