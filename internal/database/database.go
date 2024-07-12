package database

import (
	"context"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbUrl = os.Getenv("DB_URL")
)

type Database struct {
	db *mongo.Client
}

func NewDatabase() (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dbUrl).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		// return nil, err
	} else {
		log.Println("Connected to database")
	}

	log.Println("Connected to database...")

	return &Database{
		db: client,
	}, nil

}
