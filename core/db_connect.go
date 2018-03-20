package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// InitDB sdg
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=172.17.0.210 port=5432 user=controller dbname=controller password=controller sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db, nil
}
