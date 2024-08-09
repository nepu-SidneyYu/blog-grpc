package utils

import "gopkg.in/mgo.v2/bson"

func NewStringID() string {
	return bson.NewObjectId().Hex()
}
