package main

import (
	"token/models/db"
	"token/models/token"

	"github.com/sirupsen/logrus"
)

func main() {
	err := db.InitDB()
	if err != nil {
		logrus.Fatal(err)
	}

	err = db.Migrate()
	if err != nil {
		logrus.Fatal(err)
	}

	token.GenToken()
}
