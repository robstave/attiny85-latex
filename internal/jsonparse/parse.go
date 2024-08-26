package jsonparse

import (
	"encoding/json"
	"io"
	"os"
)

type Data struct {
	Title string `json:"title"`
	Color string `json:"color"`
	Fruit string `json:"fruit"`
}

func LoadJSON(jsonPath string) (Data, error) {
	var data Data

	file, err := os.Open(jsonPath)
	if err != nil {
		return data, err
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
