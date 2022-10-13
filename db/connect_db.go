package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnection *mongo.Client = connectDB()

// connectDB Connects to MongoDB Database, returns the client connection
func connectDB() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("ENV not found")
	}

	uri := os.Getenv("MONGOURI")

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connected...")
	return client
}

// GetItemDB DBGetItem Should just return the collection but its easy enough to write everything here
// Grabs an item and returns it as a Map
func GetItemDB(dbConnection *mongo.Client) (primitive.M, error) {
	coll := dbConnection.Database("craftable-items").Collection("items")
	pipeline := []primitive.D{{{Key: "$sample", Value: primitive.D{{Key: "size", Value: 1}}}}}
	cursor, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return primitive.M{}, err
	}

	var item primitive.D
	for cursor.Next(context.TODO()) {
		err = cursor.Decode(&item)
		if err != nil {
			return primitive.M{}, err
		}
	}

	itemMap := item.Map()

	//Check that is craftable item
	for itemMap["type"] == "minecraft:smelting" || itemMap["type"] == "minecraft:smithing" {
		itemMap, _ = GetItemDB(dbConnection)
	}

	return itemMap, nil
}

// Should get the 27 items that will go inside the inventory, within those items we need to ensure we have the ingredients/key items needed to craft the CurrentItem
// func GetRandomItems(dbConnection *mongo.Client) ([]primitive.M, error) {
// 	coll := dbConnection.Database("craftable-items").Collection("items")
// 	pipeline := []primitive.D{{{Key: "$sample", Value: primitive.D{{Key: "size", Value: 27}}}}}
// 	cursor, err := coll.Aggregate(context.TODO(), pipeline)
// 	if err != nil {
// 		return []primitive.M{}, err
// 	}

// 	var item primitive.D
// 	var invItems []primitive.M
// 	for cursor.Next(context.TODO()) {
// 		err = cursor.Decode(&item)
// 		invItems = append(invItems, item.Map())
// 		if err != nil {
// 			return []primitive.M{}, err
// 		}

// 	}

// 	return nil, nil
// }

func GetNumOfDocuments(dbConnection *mongo.Client) int64 {
	coll := dbConnection.Database("MinecraftItems").Collection("Items")
	estCount, err := coll.EstimatedDocumentCount(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return estCount
}
