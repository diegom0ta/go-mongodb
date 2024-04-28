package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set MongoDB connection string
	uri := "mongodb://mongodb:mongodb@127.0.0.1:27017/auth_microservice_db?authSource=admin"

	// Context with timeout for connecting to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// Ping MongoDB to check if the connection was successful
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Perform MongoDB operations here...
	// Example: Create a new database and collection
	database := client.Database("mydatabase")
	collection := database.Collection("mycollection")

	// Insert a document
	_, err = collection.InsertOne(ctx, map[string]interface{}{
		"name":  "John Doe",
		"email": "johndoe@example.com",
	})
	if err != nil {
		log.Fatal("Failed to insert document:", err)
	}

	fmt.Println("Document inserted successfully!")
}
