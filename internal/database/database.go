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
	user  = os.Getenv("DB_USERNAME")
	pass  = os.Getenv("DB_ROOT_PASSWORD")
	addr  = os.Getenv("DB_ADDR")
	dbUrl = os.Getenv("DB_URL")
)

type Database struct {
	db *mongo.Client
}

func NewDatabase() (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	} else {
		log.Println("Connected to database")
	}

	return &Database{
		db: client,
	}, nil

}
