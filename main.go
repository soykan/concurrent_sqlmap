package main

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string
	Description string
}

type Request struct {
	gorm.Model
	Raw       string
	ProjectID int
	Project   Project
}

type Parameter struct {
	gorm.Model
	Name      string
	RequestID int
	Request   Request
}

func main() {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting the database.")
	}
	db.AutoMigrate(&Project{}, &Request{}, &Parameter{})

}
