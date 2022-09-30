package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	UsersList []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	readFile()
}

func readFile() {
	// Open our jsonFile
	jsonFile, err := os.Open("users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users Users

	json.Unmarshal(byteValue, &users)

	for _, user := range users.UsersList {
		fmt.Printf("User Name: %s\n", user.Name)
		fmt.Printf("Type: %s\n", user.Type)
		fmt.Printf("Age: %v\n", user.Age)
		fmt.Printf("Facebook: %s\n", user.Social.Facebook)
		fmt.Printf("Twitter: %s\n", user.Social.Twitter)
		fmt.Println()
	}
}
