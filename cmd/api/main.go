package main

import (
	"fmt"

	transportHTTP "MailBeacon/internal/transport/http"
)

func Run() error {
	fmt.Println("Starting server...")
	// todo: set up db
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
