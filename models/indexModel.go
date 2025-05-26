package models

import (
	"bytes"
	"embed"
	"encoding/json"
	"waifu_gcx/entities"
)

func LoadJson(data embed.FS) ([]entities.Waifu, error) {
	file, err := data.ReadFile("data/waifu.json")
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(file)
	var result []entities.Waifu
	err = json.NewDecoder(reader).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result, err

}
