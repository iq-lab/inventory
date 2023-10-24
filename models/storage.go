/*
Each unique item I can store has a StorageItemType. I could have granola bars in my pantry and my backpack, and in this case, each would have the same StorageItemType,
but each would have their own StorageItem in the database. On top of this, there are ParentStorages, which are things like the fridge, my backpack, the pantry, etc.
These have things called child storages, which would be which pocket in my backpack or which shelf of the fridge, etc.

Each ChildStorage has its own location code, which can be anything I want, like fridge:shelf1, or something like that.
*/

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StorageItemType struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	UOM  string `json:"uom"`
}

type StorageItem struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	Qty      string `json:"qty"`
	Location string `json:"location"`
}

type ParentStorage struct {
	Name string             `json:"name" bson:"name"`
	ID   primitive.ObjectID `json:"id" bson:"_id"`
}

type ChildStorage struct {
	ParentID     string `json:"parentId"`
	LocationCode string `json:"locCode"`
}
