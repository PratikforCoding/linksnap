package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() {

	
	// fetching the connection link from the environment variable
	mongoUri := os.Getenv("DB_CONNECTION_LINK")
	// Setting server api version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoUri).SetServerAPIOptions(serverAPI)

	// Create new client to connect to server
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}
	
	DB = client

	// Ping the deployment to check if it's connected
	var result bson.M
	if err := DB.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode((&result)); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

}

// GetCollection function to get the collection from the database
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database("linksnap").Collection(collectionName)
}

// DisconnectDB function to disconnect from the database
func DisconnectDB() {
	if err := DB.Disconnect(context.TODO()) ; err != nil {
		panic(err)
	}
}