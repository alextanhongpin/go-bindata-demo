package main

import (
	"encoding/json"
	"log"

	"github.com/alextanhongpin/go-binddata-demo/asset"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//go:generate go-bindata -o asset/bindata.go data/
func main() {
	log.Println("hello world")

	data, err := asset.Asset("data/data.json")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(data))
	var people []People
	err = json.Unmarshal(data, &people)
	if err != nil {
		log.Println(err)
	}
	log.Println(people)
}
