package core

import (
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
)

// InitDB sdg
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic(err)
	}
	return db, nil
}
