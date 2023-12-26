package main

import (
	"encoding/json"
	"errors"
	"os"
)

const CONFIG_PATH = "../configs/server.json"

type Config struct {
	Save_path    string
	History_size int
}

func ReadConfig() (Config, error) {
	var payload Config

	content, err := os.ReadFile(CONFIG_PATH)

	if err != nil {
		return payload, errors.New("Error opening config file.")
	}

	err = json.Unmarshal(content, &payload)
	if err != nil {
		return payload, errors.New("Error during Unmarshal.")
	}

	return payload, nil
}
