package controllers

import (
	"context"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/allefts/minedle/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemsUsed struct {
	Name string
	URL  string
}

var ItemsUsedInCrafting []ItemsUsed

// Should get the 27 items that will go inside the inventory, within those items we need to ensure we have the ingredients/key items needed to craft the CurrentItem
func GetRandomItems() ([]ItemsUsed, error) {
	//empty it out
	ItemsUsedInCrafting = ItemsUsedInCrafting[:0]

	//connect do collection
	coll := db.MongoConnection.Database("craftable-items").Collection("items")
	currItem := CurrentItem

	itemUsed := ItemsUsed{}
	itemName := ""

	if currItem["type"] == "minecraft:crafting_shapeless" {
		//Look for ingredients
		for _, val1 := range currItem["ingredients"].(primitive.A) {
			for _, val2 := range val1.(primitive.D) {
				itemName = strings.ReplaceAll(val2.Value.(string), "minecraft:", "")
				itemUsed.Name = itemName
				itemUsed.URL = "https://minedle-pictures.s3.us-east-2.amazonaws.com/" + itemName + ".png"
				ItemsUsedInCrafting = append(ItemsUsedInCrafting, itemUsed)
			}
		}
	} else if currItem["type"] == "minecraft:crafting_shaped" {
		//Look for key
		for _, val1 := range currItem["key"].(primitive.D) {
			for _, val2 := range val1.Value.(primitive.D) {
				itemName = strings.ReplaceAll(val2.Value.(string), "minecraft:", "")
				itemUsed.Name = itemName
				itemUsed.URL = "https://minedle-pictures.s3.us-east-2.amazonaws.com/" + itemName + ".png"
				ItemsUsedInCrafting = append(ItemsUsedInCrafting, itemUsed)
			}
		}
	}

	numOfItemsLeft := 27 - len(ItemsUsedInCrafting)
	pipeline := []primitive.D{{{Key: "$sample", Value: primitive.D{{Key: "size", Value: numOfItemsLeft}}}}}
	cursor, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return []ItemsUsed{}, err
	}

	var item primitive.M
	// var invItems []primitive.M
	for cursor.Next(context.TODO()) {
		err = cursor.Decode(&item)
		// fmt.Println(reflect.TypeOf(item["result"]))
		if reflect.TypeOf(item["result"]).String() == "primitive.M" {
			itemName = strings.ReplaceAll((item["result"]).(primitive.M)["item"].(string), "minecraft:", "")
			itemUsed.Name = itemName
			itemUsed.URL = "https://minedle-pictures.s3.us-east-2.amazonaws.com/" + itemName + ".png"
			ItemsUsedInCrafting = append(ItemsUsedInCrafting, itemUsed)
		} else {
			itemName = strings.ReplaceAll(item["result"].(string), "minecraft:", "")
			itemUsed.Name = itemName
			itemUsed.URL = "https://minedle-pictures.s3.us-east-2.amazonaws.com/" + itemName + ".png"
			ItemsUsedInCrafting = append(ItemsUsedInCrafting, itemUsed)
		}

		if err != nil {
			return []ItemsUsed{}, err
		}
		// invItems = append(invItems, item)
	}

	//Randomize Order
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ItemsUsedInCrafting), func(i, j int) {
		ItemsUsedInCrafting[i], ItemsUsedInCrafting[j] = ItemsUsedInCrafting[j], ItemsUsedInCrafting[i]
	})

	// fmt.Println(ItemsUsedInCrafting)
	return ItemsUsedInCrafting, nil
}
