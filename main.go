package main

import (
	"context"
	"log"
	"os"

	"github.com/FromhelStudio/fromhel-clients-tables-fix/usecases"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		panic(uri)
	}

	conn, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := conn.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	usecases.Migrate(conn)
}
