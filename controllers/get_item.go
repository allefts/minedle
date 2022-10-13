package controllers

import (
	"time"

	"github.com/allefts/minedle/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var LastItemUpdateTime = time.Now()
var CurrentItem, _ = db.GetItemDB(db.MongoConnection)
var UsedItems = initUsedItemsSlice()

func HandleItem() (primitive.M, error) {
	var item primitive.M
	var err error
	//Checks for if its been 24 hours since the server has changed items
	newItem := checkTime(&LastItemUpdateTime, time.Now())
	//If it has been more than 24 hours, new item
	if newItem {
		item, err = db.GetItemDB(db.MongoConnection)
		if err != nil {
			return primitive.M{}, err
		}
		//Check for item already used and get new item if already used
		item, err = UsedItemContains(item)
		if err != nil {
			return primitive.M{}, err
		}
		//Change Current Item
		CurrentItem = item
		//Add Item to used items
		UsedItems = append(UsedItems, item["_id"].(primitive.ObjectID))
		//Has not been 24 hours
	} else {
		//Keep same item
		item = CurrentItem
	}

	// fmt.Println(CurrentItem)

	return item, nil
}

func checkTime(prevTime *time.Time, currTime time.Time) bool {
	timeDiff := prevTime.Sub(currTime)
	if timeDiff.Abs().Hours() > 1 {
		*prevTime = time.Now()
		return true
	}
	return false
}

func initUsedItemsSlice() []primitive.ObjectID {
	return []primitive.ObjectID{CurrentItem["_id"].(primitive.ObjectID)}
}

// Not tested but theory should work
func UsedItemContains(currItem primitive.M) (primitive.M, error) {
	newItem := currItem

	for _, id := range UsedItems {
		if id == currItem["_id"].(primitive.ObjectID) {
			newItem, err := db.GetItemDB(db.MongoConnection)
			if err != nil {
				return primitive.M{}, err
			}
			UsedItemContains(newItem)
		}
	}
	return newItem, nil
}
