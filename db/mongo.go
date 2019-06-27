package db

import (
	"os"

	"gopkg.in/mgo.v2"
)

// Connect - connecting to database
func Connect() (*mgo.Session, error) {
	session, err := mgo.Dial(os.Getenv("db_host"))

	if err != nil {
		return nil, err
	}

	return session, nil
}
