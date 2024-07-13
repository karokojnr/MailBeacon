package database

import (
	"context"
	"fmt"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		dbUrl,
	))

	if err != nil {
		fmt.Println("Error connecting to the database")
		panic(err)
	}
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	err = client.Ping(ctx, nil)

	if err != nil {
		fmt.Println("Error connecting to the database")
		panic(err)
	}

	fmt.Println("Connected to the database!")

	return &Database{
		db: client,
	}, nil

}
