package main

import (
	"fmt"

	"MailBeacon/internal/database"
	transportHTTP "MailBeacon/internal/transport/http"
)

func Run() error {
	fmt.Println("Starting server...")

	_, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
	}

	// todo: add more services here and add db to them

	httpHandler := transportHTTP.NewHandler()
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
