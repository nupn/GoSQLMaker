package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Page struct {
	ID	int	`json:"id"`
	Title	string	`json:"title"`
	Url	string	`json:"url"`
}

type RootA struct {
	AA []Page `json:"AA"`
}


func (p RootA) toString() string {
	return toJson(p)
}


func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

func getPages() RootA {
	raw, err := ioutil.ReadFile("./pages.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c RootA
	json.Unmarshal(raw, &c)
	return c
}

func main() {
	p := getPages()
	fmt.Println(toJson(p))
}