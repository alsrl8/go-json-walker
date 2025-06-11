package parser

import (
	"encoding/json"
	"os"
)

func ReadJsonFile(filename string) (map[string]interface{}, error) {
	return readJsonFile(filename)
}

func readJsonFile(filename string) (map[string]interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var result map[string]interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func SearchJsonFile(filename string, key string) (interface{}, error) {
	return searchJsonFile(filename, key)
}

func searchJsonFile(filename string, key string) (interface{}, error) {
	data, err := readJsonFile(filename)
	if err != nil {
		return nil, err
	}

	if value, exists := data[key]; exists {
		return value, nil
	}

	return nil, nil
}
