package models

import (
	"fmt"
	"os"
	"time"

	"bitbucket.org/joisandresky/qurban-v2/db"
	"gopkg.in/mgo.v2/bson"
)

var dbName = os.Getenv("db_name")

// Mitra - Model
type Mitra struct {
	ID      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name    string        `json:"name" bson:"name"`
	Pob     string        `json:"pob" bson:"pob"`
	Dob     time.Time     `json:"dob" bson:"dob"`
	Address Address       `json:"address" bson:"address"`
	User    bson.ObjectId `json:"user" bson:"user"`
}

// Address - address model for Mitra address
type Address struct {
	Personal string `json:"personal" bson:"personal"`
	Business string `json:"business" bson:"businesss"`
}

// GetMitra - get all mitra
func (mitra Mitra) GetMitra() ([]Mitra, error) {
	var mitras []Mitra
	session, err := db.Connect()
	if err != nil {
		return mitras, err
	}

	if err = session.DB(dbName).C("mitra").Find(bson.M{}).All(&mitras); err != nil {
		return mitras, err
	}

	return mitras, err
}

// ShowMitra - showing one mitra
func (mitra *Mitra) ShowMitra() (Mitra, error) {
	var foundMitra Mitra
	session, err := db.Connect()
	if err != nil {
		return foundMitra, err
	}

	if err = session.DB(dbName).C("mitra").Find(bson.M{"_id": mitra.ID}).One(&foundMitra); err != nil {
		return foundMitra, err
	}

	return foundMitra, err
}

// InsertMitra - insert One Mitra and User model
func (mitra *Mitra) InsertMitra(userData User) error {
	session, err := db.Connect()
	if err != nil {
		return err
	}

	if err = session.DB(dbName).C("user").Insert(&userData); err != nil {
		fmt.Println("Err to create User")
		return err
	}

	if err = session.DB(dbName).C("mitra").Insert(&mitra); err != nil {
		fmt.Println("Err to create Mitra")
		return err
	}

	return nil
}

// UpdateMitra - update one mitra
func (mitra *Mitra) UpdateMitra(id string) error {
	session, err := db.Connect()
	if err != nil {
		return err
	}

	if err = session.DB(dbName).C("mitra").Update(bson.M{"_id": bson.ObjectIdHex(id)}, mitra); err != nil {
		return err
	}

	return nil
}

// RemoveMitra - remove one mitra
func (mitra *Mitra) RemoveMitra(id string) error {
	session, err := db.Connect()
	if err != nil {
		return err
	}

	if err = session.DB(dbName).C("mitra").Remove(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
		return err
	}

	return nil
}
