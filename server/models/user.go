package models

import (
	"encoding/json"
	"os"
)

type User struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

const (
	jsonFile = "user.json"
)

var listUsers []User
var id = map[string]int{}

func Init() {
	id = make(map[string]int)

}
func IsExisted(userName string) bool {
	_, existed := id[userName]
	return existed
}

func Register(userName string, password string) bool {
	if IsExisted(userName) {
		return false
	}
	id[userName] = len(listUsers)
	listUsers = append(listUsers, User{Username: userName, Password: password})
	return true
}

func UpdateInfo(userName string, password string) bool {
	if !IsExisted(userName) {
		return false
	}
	i := id[userName]
	listUsers[i].Password = password
	return true
}

func readUserListFromFile() bool {
	data, _ := os.ReadFile(jsonFile)
	err := json.Unmarshal(data, &listUsers)
	if err != nil {
		panic(err)
		return false
	}
	return true
}
func saveUserListToFile() bool {
	data, err := json.Marshal(listUsers)
	if err != nil {
		panic(err)
		return false
	}
	os.WriteFile(jsonFile, data, 0644)
	return true
}
