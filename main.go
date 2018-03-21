package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"./core"
	"./models"
	"github.com/jinzhu/gorm"
	yaml "gopkg.in/yaml.v2"
)

var db *gorm.DB

// Config wet
type Config struct {
	Appname string `yaml:"appname"`
	Db      struct {
		Filename string `yaml:"filename"`
	}
	WebServer struct {
		Port int `yaml:"port"`
	}
	TestVal string `yaml:"testval"`
}

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var err error
	db, err = core.InitDB()
	check(err)

	yamlFile, err := ioutil.ReadFile("conf/config.yml")
	check(err)
	config := Config{}
	err = yaml.Unmarshal(yamlFile, &config)
	check(err)

	http.HandleFunc("/user", GetUser)
	listen := ":" + strconv.Itoa(config.WebServer.Port)
	http.ListenAndServe(listen, nil)
}
