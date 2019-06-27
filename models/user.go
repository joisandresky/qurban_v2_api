package models

import (
	"gopkg.in/mgo.v2/bson"
)

// User - model
type User struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password"  bson:"password"`
	Role     string        `json:"role" bson:"role"`
}
