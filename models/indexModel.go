package models

import (
	"encoding/json"
	"os"
)

type Waifu struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func LoadJson(path string) ([]Waifu, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var result []Waifu
	err = json.NewDecoder(file).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result, err

}
