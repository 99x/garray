package garrays

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
	Id       string `json:"_id"`
	Guid     string `json:"guid"`
	IsActive bool   `json:"isActive"`
	Age      int    `json:"age"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
}

var users = *ReadUserData()

func ReadUserData() *[]User {

	// Open our jsonFile
	jsonFile, err := os.Open("./test_data/users_data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Print(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users []User
	json.Unmarshal(byteValue, &users)

	return &users
}
