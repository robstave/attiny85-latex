package jsonparse

import (
	"attiny85-latex/internal/data"
	"encoding/json"
	"io"
	"os"
)

// OpenFile opens the file at the given path and returns an *os.File.
func OpenFile(jsonPath string) (*os.File, error) {
	file, err := os.Open(jsonPath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// ReadFile reads the content of the provided file and returns it as a byte slice.
func ReadFile(file *os.File) ([]byte, error) {
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return byteValue, nil
}

// ParseJSON parses the provided byte slice into a Data struct.
func ParseJSON(byteValue []byte) (data.Data, error) {
	var data data.Data
	err := json.Unmarshal(byteValue, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// LoadJSON is the original function that uses the smaller functions to load and parse JSON.
func LoadJSON(jsonPath string) (data.Data, error) {
	file, err := OpenFile(jsonPath)
	if err != nil {
		return data.Data{}, err
	}

	byteValue, err := ReadFile(file)
	if err != nil {
		return data.Data{}, err
	}

	data, err := ParseJSON(byteValue)
	if err != nil {
		return data, err
	}

	return data, nil
}
