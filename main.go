package main

import (
	"encoding/json"
	"log"
	"net/http"

	"./core"
	"./models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// GetUser Function for user
func GetUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		persons := models.Users{}
		db.Find(&persons)

		js, err := json.Marshal(persons)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	case "PUT":
		person := models.User{}
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			panic(err)
		}
		db.Save(&person)
	}

}

func main() {
	ldb, err := core.InitDB()
	if err != nil {
		log.Panic(err)
	}
	db = ldb

	http.HandleFunc("/user", GetUser)
	http.ListenAndServe(":8000", nil)

}
