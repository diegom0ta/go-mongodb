package database

import (
	"context"
	"fmt"
	"log"

	"github.com/diegom0ta/go-mongodb/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Connect to MongoDB
func Connect(ctx context.Context) {
	config, err := utils.ParseYaml()
	if err != nil {
		log.Fatalf("Error parsing yaml: %v", err)
	}

	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/auth_microservice_db?authSource=admin", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port)

	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	log.Println("Connected to MongoDB!")

}

// Disconnect from database
func Disconnect(ctx context.Context) {
	err := Client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
