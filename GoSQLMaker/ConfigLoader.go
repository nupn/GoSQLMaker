package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type GeneralInfo struct {
	SQLServer	string	`json:"SQLServer"`
	SQLPort	int	`json:"SQLPort"`
	SQLUser	string	`json:"SQLUser"`
	SQLPassWord	string	`json:"SQLPassWord"`
}

type TableInfo struct {
	SQL	string	`json:"SQL"`
}

type ConfigInfo struct {
	General GeneralInfo `json:"General"`
	Tables []TableInfo `json:"Tables"`
	TableHash []TableInfo `json:"TableHash"`	
}

func (p ConfigInfo) toString() string {
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

func GetConfig() ConfigInfo {
	raw, err := ioutil.ReadFile("./Config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c ConfigInfo
	json.Unmarshal(raw, &c)
	return c
}
